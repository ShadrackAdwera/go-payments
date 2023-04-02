package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
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
	// _, err := srv.IsAuthorized(ctx, "", utils.RequestsRead)

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
		Offset: getRequestsArgs.PageID,
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
