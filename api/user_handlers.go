package api

import (
	"net/http"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserArgs struct {
	Username string       `json:"username" binding:"required,min=5"`
	Email    string       `json:"email" binding:"required,email"`
	Role     db.UserRoles `json:"role" binding:"required"`
}

type UserResponse struct {
	Message string  `json:"message"`
	User    db.User `json:"user"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var createUserArgs CreateUserArgs

	if err := ctx.ShouldBindJSON(&createUserArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	userId, err := uuid.NewRandom()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		ID:       userId,
		Username: createUserArgs.Username,
		Email:    createUserArgs.Email,
		Role:     createUserArgs.Role,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	response := UserResponse{
		Message: "user created",
		User:    user,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (s *Server) getUser(ctx *gin.Context) {

}

func (s *Server) getUsers(ctx *gin.Context) {

}

func (s *Server) updateUser(ctx *gin.Context) {

}

func (s *Server) deleteUser(ctx *gin.Context) {

}
