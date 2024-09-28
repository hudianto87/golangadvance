package handler

import (
	"belajargolangpart2/session11-user-crud-grpc-gateway-cache/entity"
	"belajargolangpart2/session11-user-crud-grpc-gateway-cache/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = fieldErrorMessage(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdUser, err := h.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	c.JSON(http.StatusOK, createdUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updateUser, err := h.userService.UpdateUser(c.Request.Context(), id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "user deleted"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func fieldErrorMessage(errorMessage string) string {
	switch {
	case strings.Contains(errorMessage, "'Name' failed on the 'required' tag"):
		return "Name is mandatory column"
	case strings.Contains(errorMessage, "'Name' failed on the 'min' tag"):
		return "Name minimal 3 character"
	case strings.Contains(errorMessage, "'Email' failed on the 'required' tag"):
		return "Email is mandatory column"
	case strings.Contains(errorMessage, "'Email' failed on the 'email' tag"):
		return "Email not valid"
	}

	return errorMessage
}
