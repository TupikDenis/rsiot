package subject

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetAllSubjects() []models.TransformedSubject {
	var subjects []models.Subject
	var _subjects []models.TransformedSubject

	handlers.Database().Order("id asc").Find(&subjects)

	for _, item := range subjects {
		_subjects = append(
			_subjects, models.TransformedSubject{
				Id:   item.ID,
				Name: item.Name,
			})
	}

	return _subjects
}
