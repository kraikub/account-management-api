package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {}

func (handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
