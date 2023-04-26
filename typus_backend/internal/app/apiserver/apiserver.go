package apiserver

import (
	"backend/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Server        *gin.Engine
	AuthHandler   *handlers.AuthHandler
	SampleHandler *handlers.SampleHandler
	// logger *logger.Logger
}

func NewAPIServer() *APIServer {
	server := gin.Default()
	authHandler := handlers.NewAuthHandler()
	// sampleHandler := handlers.NewSampleHandler()

	api := server.Group("/api")
	{
		authHandler.Routes(api.Group("/auth"))
		// sampleHandler.Routes(api.Group("/samples"))
	}

	return &APIServer{
		Server:      server,
		AuthHandler: authHandler,
		// SampleHandler: sampleHandler,
	}
}

func (s *APIServer) Run(port string) {
	s.Server.Run(":" + port)
}
