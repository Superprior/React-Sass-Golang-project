package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/usecases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth API endpotins handler.
// Stores its usecase and implements Handler interface.
type AuthHandler struct {
	UseCase *usecases.AuthUsecase
}

// NewAuthHandler
// This function creates a new AuthHandler.
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		UseCase: usecases.NewAuthUsecase(),
	}
}

// Routes
// This functions defines all the auth API endpoints.
func (h *AuthHandler) Routes(g *gin.RouterGroup) {
	g.POST("/register/", h.handleRegister)
	g.POST("/login/", h.handleLogin)
}

// RegisterUser
// This function implements handler for the /api/auth/register/ API endpoint.
// Register a new user with unique username and email.
func (h *AuthHandler) handleRegister(ctx *gin.Context) {

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

// LoginUser
// This function implements handler for the /api/auth/login/ API endpoint.
// Log in the user with the provided credentials, generate a JWT.
func (h *AuthHandler) handleLogin(ctx *gin.Context) {

	var logBody models.LoginCredentials

	// Read the request's body
	if err := ctx.BindJSON(&logBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Could not decode the request",
		})
		fmt.Println(err.Error())
		return
	}

	// Check if all the necessary data was provided

	if logBody.Email == "" && logBody.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Email or Username was not provided",
		})
		return
	}
	if logBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Password was not provided",
		})
		return
	}

	// Call the usecase

	token, err := h.UseCase.LoginUser(logBody)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Return id of the created user

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
