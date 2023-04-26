package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct{}

func (h *SampleHandler) Routes(g *gin.RouterGroup) {
	g.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Route": "/api/samples"})
	})
	g.GET("/:sampleId", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Route": fmt.Sprintf("/api/samples/%s", ctx.Param("sampleId"))})
	})
}
