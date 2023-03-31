package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreatePermissionArgs struct {
	Name        string `json:"name" binding:"required,min=5"`
	Description string `json:"description" binding:"required,min=5"`
}

func (srv *Server) createPermission(ctx *gin.Context) {
	// p := getProfileData(ctx)
	// if p.Sub == "" {
	// 	ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
	// 	return
	// }

	var permArgs CreatePermissionArgs

	if err := ctx.ShouldBindJSON(&permArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	perm, err := srv.store.CreatePermission(ctx, db.CreatePermissionParams{
		Name:        permArgs.Name,
		Description: permArgs.Description,
		CreatedbyID: "",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": perm})
}

type PermissionParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (srv *Server) getPermissionById(ctx *gin.Context) {
	var permissionParams PermissionParams

	if err := ctx.ShouldBindUri(&permissionParams); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	permission, err := srv.store.GetPermission(ctx, permissionParams.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("this permission does not exist")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": permission})
}

type GetPermissionsArgs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (srv *Server) getPermissions(ctx *gin.Context) {
	var getPermissionsArgs GetPermissionsArgs

	if err := ctx.ShouldBindQuery(&getPermissionsArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	permissions, err := srv.store.GetPermissions(ctx, db.GetPermissionsParams{
		Limit:  getPermissionsArgs.PageSize,
		Offset: getPermissionsArgs.PageID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("this permission does not exist")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": permissions})

}
