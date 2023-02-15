package user

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func DeleteUser(id uint) bool {
	var user models.User
	var mark models.Mark

	db := handlers.Database()
	err := db.First(&user, id).Delete(&user).Error

	if err != nil {
		return false
	}

	err = db.Where("user_id = ?", id).Delete(&mark).Error

	if err != nil {
		return false
	}

	return true
}
