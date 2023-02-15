package user

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func CreateUser(c *gin.Context) bool {
	var requestBody models.UserRequestBody
	var user models.User

	err := c.BindJSON(&requestBody)

	if err != nil || requestBody.Name == "" {
		return false
	}

	user = models.User{
		Name: requestBody.Name,
	}

	db := handlers.Database()

	if err != nil {
		return false
	}

	err = db.Save(&user).Error

	if err != nil {
		return false
	}

	return true
}
