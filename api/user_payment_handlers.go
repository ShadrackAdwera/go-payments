package api

import (
	"fmt"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/gin-gonic/gin"
)

type GetUserPaymentsArgs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (s *Server) getUserPayments(ctx *gin.Context) {

	p := getProfileData(ctx)
	if p.Sub == "" {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
		return
	}

	_, err := s.IsAuthorized(ctx, p.Sub, utils.UserPaymentsRead)

	if err != nil {
		ctx.JSON(http.StatusForbidden, errJSON(err))
		return
	}

	var getUserPaymentArgs GetUserPaymentsArgs

	if err := ctx.ShouldBindUri(&getUserPaymentArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	up, err := s.store.GetUserPayments(ctx, db.GetUserPaymentsParams{
		Limit:  getUserPaymentArgs.PageSize,
		Offset: (getUserPaymentArgs.PageID - 1) * getUserPaymentArgs.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": up})

}
