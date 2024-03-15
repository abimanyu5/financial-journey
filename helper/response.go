package helper

import (
    "github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse represents a standard success response structure
type SuccessResponse struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// RespondWithError sends an error response with the given message and status code
func RespondWithError(c *gin.Context, statusCode int, message string, data interface{}) {
    response := ErrorResponse{
        Status:  statusCode,
        Message: message,
        Data:    data,
    }
    c.JSON(statusCode, response)
}

// RespondWithSuccess sends a success response with the given message and status code
func RespondWithSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
    response := SuccessResponse{
        Status:  statusCode,
        Message: message,
        Data:    data,
    }
    c.JSON(statusCode, response)
}
