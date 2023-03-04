package handlers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rsiot/pkg/models"
)

func Database() *gorm.DB {
	/*host=localhost user=gorm password=gorm dbname=gorm port=9920*/
	/*"host=34.118.127.239 user=postgres password=root dbname=postgres port=5432"*/
	dsn := "host=34.118.127.239 user=postgres password=root dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Mark{})

	return db
}
