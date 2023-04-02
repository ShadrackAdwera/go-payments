package api

import (
	"database/sql"
	"fmt"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (srv *Server) IsAuthorized(ctx *gin.Context, userId string, permissionName string) (db.UsersPermission, error) {

	perm, err := srv.store.GetPermissionByName(ctx, permissionName)

	if err != nil {
		if err == sql.ErrNoRows {
			return db.UsersPermission{}, fmt.Errorf("this permission does not exist")
		}
		return db.UsersPermission{}, fmt.Errorf("an error occured while fetching the permissions")
	}

	userPerm, err := srv.store.GetUserPermissionByUserIdAndPermissionId(ctx, db.GetUserPermissionByUserIdAndPermissionIdParams{
		UserID:       userId,
		PermissionID: perm.ID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return db.UsersPermission{}, fmt.Errorf("this user is not authorized to perform this action")
		}
		return db.UsersPermission{}, fmt.Errorf("an error occured while fetching the user permissions")
	}
	return userPerm, nil
}
