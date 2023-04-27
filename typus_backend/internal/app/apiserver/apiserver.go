package apiserver

import (
	"backend/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Server        *gin.Engine
	AuthHandler   *handlers.Handler
	SampleHandler *handlers.Handler
	// logger *logger.Logger
}

func NewAPIServer() *APIServer {
	server := gin.Default()
	var authHandler handlers.Handler = handlers.NewAuthHandler()
	var sampleHandler handlers.Handler = handlers.NewSampleHandler()

	api := server.Group("/api")
	{
		authHandler.Routes(api.Group("/auth"))
		sampleHandler.Routes(api.Group("/samples"))
	}

	return &APIServer{
		Server:        server,
		AuthHandler:   &authHandler,
		SampleHandler: &sampleHandler,
	}
}

func (s *APIServer) Run(port string) {
	s.Server.Run(":" + port)
}
