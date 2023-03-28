package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserArgs struct {
	ID string `uri:"id" binding:"required"`
}

func (s *Server) getUserById(ctx *gin.Context) {
	var user GetUserArgs

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	foundUser, err := s.store.GetUser(ctx, user.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("this user was not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, foundUser)
}
