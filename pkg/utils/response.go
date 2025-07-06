package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}
func UnAuthorizeResponse(c *gin.Context, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
