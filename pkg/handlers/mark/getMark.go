package mark

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetMarkById(id uint) models.TransformedMark {
	var mark models.Mark

	handlers.Database().First(&mark, id)

	markById := models.TransformedMark{
		Id:        mark.ID,
		UserId:    mark.Mark,
		SubjectId: mark.SubjectId,
		Mark:      mark.Mark,
	}

	return markById
}
