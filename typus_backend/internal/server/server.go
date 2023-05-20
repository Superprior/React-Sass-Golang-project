package server

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Server struct definition.
// Contains Gin engine, all handlers and the logger.
type APIServer struct {
	Server        *gin.Engine
	AuthHandler   *handlers.Handler
	SampleHandler *handlers.Handler
	// logger *logger.Logger
}

// Create a new APIServer instance.
func NewAPIServer() (s *APIServer) {
	server := gin.Default()
	var authHandler handlers.Handler = handlers.NewAuthHandler()
	var sampleHandler handlers.Handler = handlers.NewSampleHandler()

	api := server.Group("/api")
	{
		authHandler.Routes(api.Group("/auth"))
		sampleHandler.Routes(api.Group("/samples"))
	}

	s = &APIServer{
		Server:        server,
		AuthHandler:   &authHandler,
		SampleHandler: &sampleHandler,
	}
	return
}

// Run the APIServer.
func (s *APIServer) Run(port string) {
	s.Server.Run(":" + port)
}
