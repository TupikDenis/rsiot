package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name  string `json:"name"`
	Marks []Mark `gorm:"ForeignKey:SubjectId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransformedSubject struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type SubjectRequestBody struct {
	Name string
}
