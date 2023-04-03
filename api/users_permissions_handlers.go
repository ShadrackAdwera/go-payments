package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateUserPermissionArgs struct {
	UserID       string `json:"user_id" binding:"required"`
	PermissionID int64  `json:"permission_id" binding:"required,min=1"`
}

func (srv *Server) createUserPermission(ctx *gin.Context) {
	var createUserPermissioArgs CreateUserPermissionArgs

	if err := ctx.ShouldBindJSON(&createUserPermissioArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	userPerm, err := srv.store.CreateUserPermission(ctx, db.CreateUserPermissionParams{
		UserID:       createUserPermissioArgs.UserID,
		PermissionID: createUserPermissioArgs.PermissionID,
		CreatedbyID:  "",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user_permission": userPerm})
}

type GetByUserIdAndPermIdArgs struct {
	UserID       string `form:"user_id" binding:"required"`
	PermissionID int64  `form:"permission_id" binding:"required,min=1"`
}

func (srv *Server) getByUserIdAndPermId(ctx *gin.Context) {
	var getByUserIdAndPermIdArgs GetByUserIdAndPermIdArgs

	if err := ctx.ShouldBindQuery(&getByUserIdAndPermIdArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	userPerm, err := srv.store.GetUserPermissionByUserIdAndPermissionId(ctx, db.GetUserPermissionByUserIdAndPermissionIdParams{
		UserID:       getByUserIdAndPermIdArgs.UserID,
		PermissionID: getByUserIdAndPermIdArgs.PermissionID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("this permission linked to the user does not exist")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_permission": userPerm})
}
