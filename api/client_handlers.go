package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateClientArgs struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Phone                string `json:"phone" binding:"required"`
	AccountNumber        string `json:"account_number"`
	PreferredPaymentType string `json:"preferred_payment_type" binding:"required"`
}

func (s *Server) createClient(ctx *gin.Context) {

	p := getProfileData(ctx)
	if p.Sub == "" {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
		return
	}

	var createClientArgs CreateClientArgs

	if err := ctx.ShouldBindJSON(&createClientArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	client, err := s.store.CreateClient(ctx, db.CreateClientParams{
		Name:  createClientArgs.Name,
		Email: createClientArgs.Email,
		Phone: createClientArgs.Phone,
		AccountNumber: sql.NullString{
			String: createClientArgs.AccountNumber,
			Valid:  true,
		},
		PreferredPaymentType: db.PaymentTypes(createClientArgs.PreferredPaymentType),
		CreatedbyID:          p.Sub,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"client": client})
}
