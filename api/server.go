package api

import (
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

	//user routes
	router.GET("/api/users")
	router.GET("/api/users/:id")
	router.PATCH("/api/users")
	router.DELETE("/api/users/:id")
	server.router = router
	return &server
}
