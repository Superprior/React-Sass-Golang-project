package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	Routes(g *gin.RouterGroup)
}
