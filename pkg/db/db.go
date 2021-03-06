package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"product-micro/pkg/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler{
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{db}
}
