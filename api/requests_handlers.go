package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/worker"
	"github.com/gin-gonic/gin"
)

type NewRequestArgs struct {
	Title        string `json:"title" binding:"required"`
	Amount       int64  `json:"amount" binding:"required,min=100"`
	PaidToID     int64  `json:"paid_to_id" binding:"required"`
	ApprovedbyID string `json:"approvedby_id" binding:"required"`
}

func (srv *Server) createRequest(ctx *gin.Context) {
	var newRequestArgs NewRequestArgs

	if err := ctx.ShouldBindJSON(&newRequestArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	// _, err := srv.IsAuthorized(ctx, "", utils.PaymentInitiator)

	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, errJSON(err))
	// 	return
	// }

	// approver, err := srv.IsAuthorized(ctx, newRequestArgs.ApprovedbyID, utils.PaymentInitiator)

	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, errJSON(err))
	// 	return
	// }

	// request, err := srv.store.CreateRequest(ctx, db.CreateRequestParams{
	// 	Title:       newRequestArgs.Title,
	// 	Status:      db.ApprovalStatusPending,
	// 	Amount:      newRequestArgs.Amount,
	// 	PaidToID:    newRequestArgs.PaidToID,
	// 	CreatedbyID: "",
	// 	ApprovedbyID: approver.UserID,
	// })

	request, err := srv.store.CreateRequest(ctx, db.CreateRequestParams{
		Title:        newRequestArgs.Title,
		Status:       db.ApprovalStatusPending,
		Amount:       newRequestArgs.Amount,
		PaidToID:     newRequestArgs.PaidToID,
		CreatedbyID:  "",
		ApprovedbyID: "",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"request": request})
}

type GetRequestsArgs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (srv *Server) getRequests(ctx *gin.Context) {

	// p := getProfileData(ctx)
	// if p.Sub == "" {
	// 	ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
	// 	return
	// }

	// _, err := srv.IsAuthorized(ctx, p.Sub, utils.RequestsRead)

	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, errJSON(err))
	// 	return
	// }

	var getRequestsArgs GetRequestsArgs

	if err := ctx.ShouldBindQuery(&getRequestsArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	requests, err := srv.store.GetRequests(ctx, db.GetRequestsParams{
		Limit:  getRequestsArgs.PageSize,
		Offset: (getRequestsArgs.PageID - 1) * getRequestsArgs.PageSize,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("no requests found")))
			return
		}
		ctx.JSON(http.StatusNotFound, errJSON(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"requests": requests})
}

type GetRequestsToApproveArgs struct {
	Status     string `form:"status" binding:"required"`
	ApproverID string `form:"approver_id" binding:"required"`
}

func (srv *Server) getRequestsToApprove(ctx *gin.Context) {
	// p := getProfileData(ctx)
	// if p.Sub == "" {
	// 	ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
	// 	return
	// }

	// _, err := srv.IsAuthorized(ctx, p.Sub, utils.RequestsRead)

	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, errJSON(err))
	// 	return
	// }

	var getRequestsArgs GetRequestsToApproveArgs

	if err := ctx.ShouldBindQuery(&getRequestsArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	requests, err := srv.store.GetRequestsToApprove(ctx, db.GetRequestsToApproveParams{
		Status:       db.ApprovalStatus(getRequestsArgs.Status),
		ApprovedbyID: "approver", // p.Sub - fix this
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("no requests found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"requests": requests})

}

type ApproveRequestArgs struct {
	Status string `json:"status" binding:"required"`
}

type ApproveRequestUriParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (srv *Server) approveRequest(ctx *gin.Context) {
	// p := getProfileData(ctx)
	// if p.Sub == "" {
	// 	ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
	// 	return
	// }

	// _, err := srv.IsAuthorized(ctx, p.Sub, utils.RequestsApprove)

	// if err != nil {
	// 	ctx.JSON(http.StatusForbidden, errJSON(err))
	// 	return
	// }

	var approveRequestUriParams ApproveRequestUriParams

	if err := ctx.ShouldBindUri(&approveRequestUriParams); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	var approveRequestArgs ApproveRequestArgs

	if err := ctx.ShouldBindJSON(&approveRequestArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	foundReq, err := srv.store.GetRequest(ctx, approveRequestUriParams.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("no request found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	if foundReq.Status != db.ApprovalStatusPending {
		ctx.JSON(http.StatusOK, gin.H{"request": "request has already been reviewed"})
		return
	}

	// check if approved by id is same as approver id

	request, err := srv.store.ApproveRequestTx(ctx, db.ApproveRequestTxRequest{
		ID:             approveRequestUriParams.ID,
		ApprovalStatus: db.ApprovalStatus(approveRequestArgs.Status),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	if request.Request.Status == db.ApprovalStatusApproved {
		// send to a redis queue to make payment
		err = srv.distro.DistributePayment(ctx, &worker.PaymentPayload{
			ClientID:      request.Request.PaidToID,
			Amount:        request.Request.Amount,
			UserPaymentID: request.UserPayment.ID,
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errJSON(err))
			fmt.Println("error distributing the task : %w", err)
			return
		}
		// if mpesa - mpesa details
		// if bank deposit - bank details
		fmt.Printf("Request with ID %d has been %s\n", request.Request.ID, string(request.Request.Status))
	}

	ctx.JSON(http.StatusOK, gin.H{"request": request})

}
