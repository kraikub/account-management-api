package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kraikub/account-management-api/api/v1/internal/usecases"
)

func AssignRouter(r *gin.Engine, userUseCase usecases.UserUseCase) {
	h := handler{
		userUseCase: userUseCase,
	}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", h.Hello)
		v1.GET("/user/:uid", h.FindUserWithUserId)
		v1.PUT("/user/:uid", h.UpdateUserWithUserId)
	}

}
