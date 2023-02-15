package subject

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func UpdateSubject(id uint, c *gin.Context) bool {
	var subject models.Subject
	var requestBody models.SubjectRequestBody

	err := c.BindJSON(&requestBody)

	if err != nil || requestBody.Name == "" {
		return false
	}

	db := handlers.Database().First(&subject, id)
	subject.Name = requestBody.Name
	err = db.Save(&subject).Error

	if err != nil {
		return false
	}
	return true
}
