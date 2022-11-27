package controllers

import "github.com/gin-gonic/gin"

func AssignRouter(r *gin.Engine) {
	h := handler{}
	r.GET("/", h.Hello)
}
