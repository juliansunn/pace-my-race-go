package api

import (
	"api/token"
	"api/util"
	"fmt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	db "api/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new Server instance with the given store.
//
// Parameters:
// - store: a pointer to a db.Store object.
//
// Returns:
// - a pointer to the newly created Server object.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// replace the following line with the line above to use JWT instead of Paseto
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/o
func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/users/:id/logout", server.logoutUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	// middleware
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/:id", server.getUser)
	authRoutes.PATCH("/users/:id", server.updateUser)
	authRoutes.DELETE("/users/:id", server.deleteUser)
	authRoutes.GET("/users", server.listUsers)

	server.router = router
}

// Start starts the server on the specified address.
//
// Parameters:
// - address: the address to listen on.
//
// Returns:
// - error: an error if the server failed to start.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse generates an error response in the form of a gin.H map.
//
// The function takes an error as a parameter and returns a gin.H map
// with a single key "error" and the error message as its value.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
