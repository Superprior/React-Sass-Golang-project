package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func (h *AuthHandler) Routes(g *gin.RouterGroup) {
	g.POST("/register/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Route": "/api/auth/register/"})
	})
	g.POST("/login/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Route": "/api/auth/login/"})
	})
}
