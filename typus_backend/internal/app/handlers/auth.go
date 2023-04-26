package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/usecases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UseCase *usecases.AuthUsecase
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		UseCase: usecases.NewAuthUsecase(),
	}
}

func (h *AuthHandler) Routes(g *gin.RouterGroup) {
	g.POST("/register/", h.RegisterUser)
	g.POST("/login/", h.LoginUser)
}

func (h *AuthHandler) RegisterUser(ctx *gin.Context) {

	var regBody models.RegisterCredentials

	// Read the request's body
	if err := ctx.BindJSON(&regBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Could not decode the request",
		})
		fmt.Println(err.Error())
		return
	}

	// Check if all the necessary data was provided

	if regBody.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Email was not provided",
		})
		return
	}
	if regBody.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Username was not provided",
		})
		return
	}
	if regBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Password was not provided",
		})
		return
	}

	// Call the usecase

	id, err := h.UseCase.RegisterUser(regBody)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Return id of the created user

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *AuthHandler) LoginUser(ctx *gin.Context) {
	// call usecase
}
