package routers

import (
	"github.com/gin-gonic/gin"
)

type accountHandler interface {
	Hello(c *gin.Context)
}

func InitRouter(r *gin.Engine, handler accountHandler) {
	r.GET("/", handler.Hello)
}
