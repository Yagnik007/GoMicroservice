package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response represents a standardized API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an API error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Success JSON response
func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error JSON response
func Error(c *gin.Context, status int, message string, err error) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	c.JSON(status, ErrorResponse{
		Success: false,
		Message: message,
		Error:   errStr,
	})
}

// BadRequest helper
func BadRequest(c *gin.Context, err error) {
	Error(c, http.StatusBadRequest, "Bad Request", err)
}

// InternalServerError helper
func InternalServerError(c *gin.Context, err error) {
	Error(c, http.StatusInternalServerError, "Internal Server Error", err)
}
