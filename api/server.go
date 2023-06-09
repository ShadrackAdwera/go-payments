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
	// TODO: Add authentication later
	router.POST("/api/users", IsAuthenticated, server.createUser)
	router.GET("/api/users/:id", server.getUserById)
	// router.PATCH("/api/users")
	// router.DELETE("/api/users/:id")

	// client routes
	// TODO: Add authentication later on
	router.POST("/api/clients", server.createClient)
	router.GET("/api/clients", server.getClients)

	//permission routes
	// TODO: Add authentication later on
	router.GET("/api/permissions", server.getPermissions)
	router.GET("/api/permissions/:id", server.getPermissionById)
	router.POST("/api/permissions", server.createPermission)

	// users permission routes
	// TODO: Add authentication later on
	router.POST("/api/users_permissions", server.createUserPermission)
	router.GET("/api/users_permissions", server.getByUserIdAndPermId) // to use query? Tutajua

	// requests routes
	// TODO: Add authentication later on
	router.POST("/api/requests", server.createRequest)
	router.GET("/api/requests", server.getRequests)                   // /api/requests?page_id=1&page_size=10 ||
	router.GET("/api/requests/approval", server.getRequestsToApprove) // /api/requests/approval?approver_id=1&status="pending"
	router.PATCH("/api/requests/:id/approve", server.approveRequest)
	// NEXT - Approve request

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
