package models

import "gorm.io/gorm"

type Mark struct {
	gorm.Model
	UserId    uint `json:"user_id"`
	SubjectId uint `json:"subject_id"`
	Mark      uint `json:"mark"`
}

type TransformedMark struct {
	Id        uint `json:"id"`
	UserId    uint `json:"user_id"`
	SubjectId uint `json:"subject_id"`
	Mark      uint `json:"mark"`
}

type MarkRequestBody struct {
	UserId    uint `json:"user_id"`
	SubjectId uint `json:"subject_id"`
	Mark      uint `json:"mark"`
}
