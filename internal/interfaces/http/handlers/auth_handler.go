package handlers

import (
	"github.com/gin-gonic/gin"
	"go-chat-app/internal/interfaces/dto"
	"go-chat-app/internal/usecases"
	"go-chat-app/pkg/utils"
	"net/http"
)

type AuthHandler struct {
	authUseCase *usecases.AuthUseCase
}

func NewAuthHandler(authUseCase *usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// @Summary Get Access Token from Refresh Token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest tru "Refresh token request data"
// @Success 200 {object} map[string]interface{} "Refresh token successful"
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.authUseCase.RefreshToken(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Create a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request data"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 400 {object} map[string]interface{} "Registration failed"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	response, err := h.authUseCase.Register(req)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Registration failed", err)
		return
	}
	utils.CreatedResponse(c, "User registered successfully", response)
}

// @Summary Authenticate a user
// @Description Log in a user with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request data"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 401 {object} map[string]interface{} "Login failed"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	response, err := h.authUseCase.Login(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Login failed", err)
		return
	}

	utils.SuccessResponse(c, "Login successful", response)
}
