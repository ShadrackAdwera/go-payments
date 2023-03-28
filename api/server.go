package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.TxStore
}

func NewServer(store db.TxStore) *Server {
	router := gin.Default()

	server := Server{
		store: store,
	}

	// authenticated user routes
	//router.GET("/api/users")
	router.GET("/api/users/:id", server.getUserById)
	// router.PATCH("/api/users")
	// router.DELETE("/api/users/:id")
	server.router = router
	return &server
}

func errJSON(err error) gin.H {
	return gin.H{"message": fmt.Errorf(err.Error())}
}
