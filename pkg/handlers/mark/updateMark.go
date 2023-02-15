package mark

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/handlers/subject"
	"rsiot/pkg/handlers/user"
	"rsiot/pkg/models"
)

func UpdateMark(id uint, c *gin.Context) bool {
	var mark models.Mark
	var requestBody models.MarkRequestBody

	err := c.BindJSON(&requestBody)

	if err != nil || isUserStructEmpty(requestBody.UserId, requestBody.SubjectId, requestBody.Mark) ||
		(user.GetUserById(requestBody.UserId).Id == 0 || subject.GetSubjectById(requestBody.SubjectId).Id == 0) {
		return false
	}

	db := handlers.Database().First(&mark, id)
	mark.UserId = requestBody.UserId
	mark.SubjectId = requestBody.SubjectId
	mark.Mark = requestBody.Mark
	err = db.Save(&mark).Error

	if err != nil {
		return false
	}

	return true
}
