package handler

import (
	"belajargolangpart2/session4/step4/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()

	c.JSON(http.StatusOK, users)
}
