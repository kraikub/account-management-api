package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"messsage"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func handleResponse(c *gin.Context, status int, isSuccess bool, msg string, payload interface{}) {
	c.JSON(status, StandardResponse{
		Success:   isSuccess,
		Message:   msg,
		Payload:   payload,
		Timestamp: time.Now(),
	})

}
