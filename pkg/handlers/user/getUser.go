package user

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetUserById(id uint) models.TransformedUser {
	var user models.User

	handlers.Database().First(&user, id)

	userById := models.TransformedUser{
		Id:   user.ID,
		Name: user.Name,
	}

	return userById
}
