package mark

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func DeleteMark(id uint) bool {
	var mark models.Mark
	err := handlers.Database().First(&mark, id).Delete(&mark).Error

	if err != nil {
		return false
	}

	return true
}
