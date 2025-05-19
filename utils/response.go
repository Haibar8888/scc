package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(201, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func BadRequestResponse(c *gin.Context, message string, err interface{}) {
	c.JSON(400, APIResponse{
		Status:  "fail",
		Message: message,
		Error:   err,
	})
}

func UnauthorizedResponse(c *gin.Context, message string) {
	c.JSON(401, APIResponse{
		Status:  "fail",
		Message: message,
	})
}

func ForbiddenResponse(c *gin.Context, message string) {
	c.JSON(403, APIResponse{
		Status:  "fail",
		Message: message,
	})
}

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(404, APIResponse{
		Status:  "fail",
		Message: message,
	})
}

func InternalServerErrorResponse(c *gin.Context, message string, err interface{}) {
	c.JSON(500, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
