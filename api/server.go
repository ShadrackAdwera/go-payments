package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	router *gin.Engine
	store  db.TxStore
}

func NewServer(store db.TxStore) *Server {
	server := Server{
		store: store,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validRole)
	}

	router.POST("/api/users", server.createUser)
	router.GET("/api/users", server.getUsers)
	router.GET("/api/users/:id", server.getUser)
	router.PATCH("/api/users/:id", server.updateUser)
	router.DELETE("/api/users/:id", server.deleteUser)

	server.router = router
	return &server
}

func errJSON(err error) gin.H {
	return gin.H{"message": fmt.Errorf(err.Error())}
}
