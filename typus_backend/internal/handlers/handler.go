package handlers

import (
	"backend/pkg/headers"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Routes(g *gin.RouterGroup)
}

func handleOptions(ctx *gin.Context) {
	headers.DefaultHeaders(ctx, "OPTIONS")
}
