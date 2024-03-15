package helper

import (

	"github.com/gin-gonic/gin"
)
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}


func RespondWithError(c *gin.Context, statusCode int, message string, data interface{}) {
	response := ErrorResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}


func RespondWithSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}