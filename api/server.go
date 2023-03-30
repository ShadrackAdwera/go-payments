package api

import (
	"encoding/gob"

	"github.com/ShadrackAdwera/go-payments/api/callback"
	"github.com/ShadrackAdwera/go-payments/api/login"
	"github.com/ShadrackAdwera/go-payments/api/logout"
	"github.com/ShadrackAdwera/go-payments/authenticator"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.TxStore
	auth   *authenticator.Authenticator
}

func NewServer(store db.TxStore, auth *authenticator.Authenticator) *Server {
	router := gin.Default()
	server := Server{
		store: store,
	}

	gob.Register(map[string]interface{}{})
	cookieStore := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("auth-session", cookieStore))

	router.GET("/", server.home)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)

	// authenticated user routes
	router.GET("/api/users/:id", IsAuthenticated, server.getUserById)
	router.POST("/api/users", IsAuthenticated, server.createUser)
	// router.PATCH("/api/users")
	// router.DELETE("/api/users/:id")

	// client routes
	router.POST("/api/clients", IsAuthenticated, server.createClient)

	server.router = router
	server.auth = auth
	return &server
}

func errJSON(err error) gin.H {
	return gin.H{"message": err.Error()}
}

func (s *Server) StartServer(addr string) error {
	return s.router.Run(addr)
}
