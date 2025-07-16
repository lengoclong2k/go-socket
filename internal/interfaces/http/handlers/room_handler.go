package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-chat-app/internal/interfaces/dto"
	"go-chat-app/internal/usecases"
	"go-chat-app/pkg/utils"
	"net/http"
)

type RoomHandler struct {
	roomUseCase *usecases.RoomUseCase
}

func NewRoomHandler(roomUseCase *usecases.RoomUseCase) *RoomHandler {
	return &RoomHandler{
		roomUseCase: roomUseCase,
	}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	userId := c.MustGet("user_id").(uuid.UUID)

	var req dto.CreateRoomRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	room, err := h.roomUseCase.CreateRoom(req, userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create room", err)
		return
	}

	utils.CreatedResponse(c, "Room created successfully", room)
}

func (h *RoomHandler) GetUserRooms(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	rooms, err := h.roomUseCase.GetUserRooms(userID)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to get rooms", err)
		return
	}
	utils.SuccessResponse(c, "Rooms retrieved successfully", rooms)
}

func (h *RoomHandler) JoinRoom(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	roomIDStr := c.Param("id")

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err)
		return
	}

	if err := h.roomUseCase.JoinRoom(roomID, userID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to join room", err)
		return
	}

	utils.SuccessResponse(c, "Joined room successfully", nil)
}

func (h *RoomHandler) LeaveRoom(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	roomIDStr := c.Param("id")

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err)
		return
	}

	if err := h.roomUseCase.LeaveRoom(roomID, userID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to leave room", err)
		return
	}

	utils.SuccessResponse(c, "Left room successfully", nil)
}
