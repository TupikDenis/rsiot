package subject

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func CreateSubject(c *gin.Context) bool {
	var requestBody models.SubjectRequestBody
	var subject models.Subject

	err := c.BindJSON(&requestBody)

	if err != nil || requestBody.Name == "" {
		return false
	}

	subject = models.Subject{
		Name: requestBody.Name,
	}

	db := handlers.Database()

	if err != nil {
		return false
	}

	err = db.Save(&subject).Error

	if err != nil {
		return false
	}

	return true
}
