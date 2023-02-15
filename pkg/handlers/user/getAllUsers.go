package user

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetAllUsers() []models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser

	handlers.Database().Order("id asc").Find(&users)

	for _, item := range users {
		_users = append(
			_users, models.TransformedUser{
				Id:   item.ID,
				Name: item.Name,
			})
	}

	return _users
}
