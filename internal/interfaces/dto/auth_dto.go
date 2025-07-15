package dto

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	IsOnline bool      `json:"is_online"`
}

type PublicUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IsOnline bool      `json:"is_online"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

func ToUserResponse(user *entities.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsOnline: user.IsOnline,
	}
}

func ToPublicUserResponse(user *entities.User) PublicUserResponse {
	return PublicUserResponse{
		ID:       user.ID,
		Username: user.Username,
		IsOnline: user.IsOnline,
	}
}

func ToPublicUserResponses(users []entities.User) []PublicUserResponse {
	responses := make([]PublicUserResponse, len(users))
	for i, user := range users {
		responses[i] = ToPublicUserResponse(&user)
	}
	return responses
}

func ToAuthResponse(user *entities.User, accessToken, refreshToken string) *AuthResponse {
	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         ToUserResponse(user),
	}
}
