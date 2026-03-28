package Handler

import (
	"net/http"

	"Project/Dto"
	"Project/Model"
	"Project/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input Dto.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := Model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
