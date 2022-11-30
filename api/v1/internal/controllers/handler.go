package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kraikub/account-management-api/api/v1/internal/dtos"
	"github.com/kraikub/account-management-api/api/v1/internal/usecases"
)

type handler struct {
	userUseCase usecases.UserUseCase
}

func (handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (h handler) FindUserWithUserId(c *gin.Context) {
	uid := c.Param("uid")

	// Business use cases
	user, err := h.userUseCase.FindUserWithUserId(c, uid)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, false, "FindUserWithUserId failed", nil)
	}
	handleResponse(c, http.StatusOK, true, "", user)
}

func (h handler) UpdateUserWithUserId(c *gin.Context) {
	uid := c.Param("uid")
	var uDTO dtos.UserDTO
	if err := c.ShouldBind(&uDTO); err != nil {
		handleResponse(c, http.StatusUnprocessableEntity, false, "Unprocessable Entity", nil)
		return
	}
	if err := h.userUseCase.UpdateUserWithUserId(c, uid, uDTO); err != nil {
		handleResponse(c, http.StatusInternalServerError, false, "UpdateUserWithUserId failed", nil)
		log.Fatal(err)
		return
	}
	handleResponse(c, http.StatusOK, true, "", nil)
}
