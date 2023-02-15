package user

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func UpdateUser(id uint, c *gin.Context) bool {
	var user models.User
	var requestBody models.UserRequestBody

	err := c.BindJSON(&requestBody)

	if err != nil || requestBody.Name == "" {
		return false
	}

	db := handlers.Database().First(&user, id)
	user.Name = requestBody.Name
	err = db.Save(&user).Error

	if err != nil {
		return false
	}

	return true
}
