package usecases

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"go-chat-app/internal/interfaces/dto"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUseCase struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

type AuthResponse struct {
	Token string        `json:"token"`
	User  entities.User `json:"user"`
}

func NewAuthUseCase(userRepo repositories.UserRepository, jwtSecret string) *AuthUseCase {
	return &AuthUseCase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (uc *AuthUseCase) Register(req dto.RegisterRequest) (*AuthResponse, error) {

	//check if user exists
	existingUser, _ := uc.userRepo.GetByEmail(req.Email)

	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	existingUser, _ = uc.userRepo.GetByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	//Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:        uuid.New(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		IsOnline:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	//Generate token
	token, err := uc.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (uc *AuthUseCase) Login(req dto.LoginRequest) (*AuthResponse, error) {
	user, err := uc.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Update online status
	user.IsOnline = true
	err = uc.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	token, _ := uc.generateToken(user.ID)

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (uc *AuthUseCase) generateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte(uc.jwtSecret))
}

func (uc *AuthUseCase) ValidateToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, errors.New("invalid user ID in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	return userID, nil
}
