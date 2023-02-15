package handlers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rsiot/pkg/models"
)

func Database() *gorm.DB {
	dsn := "host=localhost user=postgres password=t7208588 dbname=rsiot port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Mark{})

	return db
}
