package handlers

import (
	"backend/internal/app/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
	UseCase *usecases.SampleUsecase
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{
		UseCase: usecases.NewSampleUsecase(),
	}
}

func (h *SampleHandler) Routes(g *gin.RouterGroup) {
	g.GET("", h.handleSamplesList)
	g.GET("/:sampleId", h.handleGetSample)
}

func (h *SampleHandler) handleSamplesList(ctx *gin.Context) {

	// No need to check the JWT in this handler

	// items, err := h.UseCase.GetAllSamples()

	// if err != nil {
	// errors
	// }
}

func (h *SampleHandler) handleGetSample(ctx *gin.Context) {

	sampleIdString := ctx.Param("sampleId")
	sampleId, err := strconv.Atoi(sampleIdString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": fmt.Sprintf("Invalid sample ID: %s", sampleIdString),
		})
		return
	}

	sample, err := h.UseCase.GetSampleById(sampleId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, sample)
}
