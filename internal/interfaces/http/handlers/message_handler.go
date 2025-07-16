package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-chat-app/internal/interfaces/dto"
	"go-chat-app/internal/usecases"
	"go-chat-app/pkg/utils"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	messageUseCase *usecases.MessageUseCase
}

func NewMessageHandler(messageUseCase *usecases.MessageUseCase) *MessageHandler {
	return &MessageHandler{
		messageUseCase: messageUseCase,
	}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	var req dto.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	message, err := h.messageUseCase.SendMessage(req, userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to send message", err)
		return
	}

	utils.CreatedResponse(c, "Message sent successfully", message)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	roomIDStr := c.Query("room_id")
	if roomIDStr == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Room ID is required", nil)
		return
	}

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid room ID", err)
		return
	}

	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	if limit <= 0 || limit > 100 {
		limit = 50
	}

	messages, err := h.messageUseCase.GetRoomMessages(roomID, userID, limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to get messages", err)
		return
	}

	utils.SuccessResponse(c, "Messages retrieved successfully", messages)
}
