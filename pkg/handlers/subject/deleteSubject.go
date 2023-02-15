package subject

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func DeleteSubject(id uint) bool {
	var subject models.Subject
	var mark models.Mark
	db := handlers.Database()
	err := db.First(&subject, id).Delete(&subject).Error

	if err != nil {
		return false
	}

	err = db.Where("subject_id = ?", id).Delete(&mark).Error

	if err != nil {
		return false
	}

	return true
}
