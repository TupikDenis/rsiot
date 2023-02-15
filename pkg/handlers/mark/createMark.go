package mark

import (
	"github.com/gin-gonic/gin"
	"rsiot/pkg/handlers"
	"rsiot/pkg/handlers/subject"
	"rsiot/pkg/handlers/user"
	"rsiot/pkg/models"
)

func CreateMark(c *gin.Context) bool {
	var requestBody models.MarkRequestBody
	var mark models.Mark

	err := c.BindJSON(&requestBody)

	if err != nil || isUserStructEmpty(requestBody.UserId, requestBody.SubjectId, requestBody.Mark) ||
		(user.GetUserById(requestBody.UserId).Id == 0 || subject.GetSubjectById(requestBody.SubjectId).Id == 0) {
		return false
	}

	mark = models.Mark{
		UserId:    requestBody.UserId,
		SubjectId: requestBody.SubjectId,
		Mark:      requestBody.Mark,
	}

	db := handlers.Database()

	if err != nil {
		return false
	}

	err = db.Save(&mark).Error

	if err != nil {
		return false
	}

	return true
}

func isUserStructEmpty(userId uint, subjectId uint, mark uint) bool {
	return userId == 0 || subjectId == 0 || mark == 0
}
