package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountHandler struct{}

func NewAccountHandler() accountHandler {
	return accountHandler{}
}

func (accountHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
