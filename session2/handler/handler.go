package handler

import (
	"belajargolangpart2/session2/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int
)

// get user by id
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //konversi ke int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})

}

// get all user
func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// create user
func CreatedUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	nextID++
	user.ID = nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// update user by id
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //konversi ke int

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, u := range users {
		updateUser := entity.User{
			ID:        id,
			Name:      user.Name,
			Email:     user.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
			UpdatedAt: time.Now(),
		}

		users[i] = updateUser
		c.JSON(http.StatusOK, updateUser)
	}
}

// delete user by id
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:1], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}
