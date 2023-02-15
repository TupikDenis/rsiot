package models

import "gorm.io/gorm"

// Account model info
// @Description User account information
// @Description with user name and username
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Marks []Mark `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransformedUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UserRequestBody struct {
	Name string
}
