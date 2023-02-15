package subject

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetSubjectById(id uint) models.TransformedSubject {
	var subject models.Subject

	handlers.Database().First(&subject, id)

	subjectById := models.TransformedSubject{
		Id:   subject.ID,
		Name: subject.Name,
	}

	return subjectById
}
