package handler

import "github.com/gin-gonic/gin"

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello word",
	})
}

func PostHandler(c *gin.Context) {
	var json struct {
		Message  string `json:"message"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200,
			gin.H{
				"messae":   json.Message,
				"location": json.Location,
			})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
