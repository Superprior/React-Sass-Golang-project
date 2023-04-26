package apiserver

import (
	"backend/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Server *gin.Engine
	// logger *logger.Logger
}

func NewAPIServer() *APIServer {
	server := gin.Default()
	authHandler := &handlers.AuthHandler{}
	sampleHandler := &handlers.SampleHandler{}

	api := server.Group("/api")
	{
		authHandler.Routes(api.Group("/auth"))
		sampleHandler.Routes(api.Group("/samples"))
	}

	return &APIServer{
		Server: server,
	}
}

func (s *APIServer) Run(port string) {
	s.Server.Run(":" + port)
}
