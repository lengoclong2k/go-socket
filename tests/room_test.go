package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/interfaces/dto"
	"go-chat-app/internal/interfaces/http/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockRoomUseCase struct {
	mock.Mock
}

func (m *MockRoomUseCase) CreateRoom(req dto.CreateRoomRequest, createdBy uuid.UUID) (*entities.Room, error) {
	args := m.Called(req, createdBy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Room), args.Error(1)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    dto.CreateRoomRequest
		userID         uuid.UUID
		mockReturn     *entities.Room
		mockError      error
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "Success - Create room",
			requestBody: dto.CreateRoomRequest{
				Name:        "Test Room",
				Description: "Test Description",
				IsPrivate:   false,
			},
			userID: uuid.New(),
			mockReturn: &entities.Room{
				ID:          uuid.New(),
				Name:        "Test Room",
				Description: "Test Description",
				IsPrivate:   false,
				CreatedBy:   uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name: "Error - UseCase returns error",
			requestBody: dto.CreateRoomRequest{
				Name:        "Test Room",
				Description: "Test Description",
				IsPrivate:   false,
			},
			userID:         uuid.New(),
			mockReturn:     nil,
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "Error - Invalid request (missing name)",
			requestBody: dto.CreateRoomRequest{
				Description: "Test Description",
				IsPrivate:   false,
			},
			userID:         uuid.New(),
			mockReturn:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockUseCase := new(MockRoomUseCase)

			// Create handler with mock usecase
			handler := &handlers.RoomHandler{
				roomUseCase: mockUseCase,
			}

			if tt.requestBody.Name != "" { // Only set up mock if request is valid
				mockUseCase.On("CreateRoom", tt.requestBody, tt.userID).Return(tt.mockReturn, tt.mockError)
			}

			// Create request
			jsonBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/rooms", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Setup router
			router := setupTestRouter()
			router.Use(func(c *gin.Context) {
				c.Set("user_id", tt.userID)
				c.Next()
			})
			router.POST("/rooms", handler.CreateRoom)

			// Execute
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			if tt.expectedError {
				assert.Equal(t, false, response["success"])
			} else {
				assert.Equal(t, true, response["success"])
			}

			if tt.requestBody.Name != "" {
				mockUseCase.AssertExpectations(t)
			}
		})
	}
}
