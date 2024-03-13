package database

import (
	"fmt"
	"gading-go-challenge/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {

	err := db.AutoMigrate(&models.Item{}, &models.Order{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Migrated")

}