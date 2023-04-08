package api

import (
	"encoding/gob"
	"time"

	"github.com/ShadrackAdwera/go-payments/api/callback"
	"github.com/ShadrackAdwera/go-payments/api/login"
	"github.com/ShadrackAdwera/go-payments/api/logout"
	"github.com/ShadrackAdwera/go-payments/authenticator"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/worker"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.TxStore
	auth   *authenticator.Authenticator
	distro worker.TaskDistributor
}

func NewServer(store db.TxStore, auth *authenticator.Authenticator, distro worker.TaskDistributor) *Server {
	router := gin.Default()
	server := Server{
		store:  store,
		distro: distro,
	}

	gob.Register(map[string]interface{}{})
	cookieStore := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("auth-session", cookieStore))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", server.home)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)

	// authenticated user routes
	// TODO: Add authentication later
	router.POST("/api/users", IsAuthenticated, server.createUser)
	router.GET("/api/users/:id", IsAuthenticated, server.getUserById)
	router.GET("/api/users", IsAuthenticated, server.getUsers)
	router.GET("/api/users/:id/permissions", IsAuthenticated, server.getPermissionsByUserId)
	// router.PATCH("/api/users")
	// router.DELETE("/api/users/:id")

	// client routes
	// TODO: Add authentication later on
	router.POST("/api/clients", IsAuthenticated, server.createClient)
	router.GET("/api/clients", IsAuthenticated, server.getClients)

	//permission routes
	// TODO: Add authentication later on
	router.GET("/api/permissions", IsAuthenticated, server.getPermissions)
	router.GET("/api/permissions/:id", IsAuthenticated, server.getPermissionById)
	router.POST("/api/permissions", IsAuthenticated, server.createPermission)
	router.POST("/api/user-permissions", IsAuthenticated, server.addPermissionsToUser)

	// users permission routes
	// TODO: Add authentication later on
	router.POST("/api/users_permissions", server.createUserPermission)
	router.GET("/api/users_permissions", server.getByUserIdAndPermId) // to use query? Tutajua

	// requests routes
	// TODO: Add authentication later on
	router.POST("/api/requests", IsAuthenticated, server.createRequest)
	router.GET("/api/requests", IsAuthenticated, server.getRequests)                   // /api/requests?page_id=1&page_size=10 ||
	router.GET("/api/requests/approval", IsAuthenticated, server.getRequestsToApprove) // /api/requests/approval?approver_id=1&status="pending"
	router.PATCH("/api/requests/:id/approve", IsAuthenticated, server.approveRequest)
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

/*
SELECT * FROM users_permissions JOIN users ON users_permissions.user_id = users.id JOIN permissions ON users_permissions.permission_id = permissions.id WHERE user_id = auth0|64186f3b8cca2db234b4f009
*/
