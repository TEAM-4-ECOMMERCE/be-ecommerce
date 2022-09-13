package utils

import (
	"e-commerce/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	categoryModel "e-commerce/domains/category/model"
	productmodel "e-commerce/domains/product/models"
	usersmodel "e-commerce/domains/users/models"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(categoryModel.Category),
		new(usersmodel.User),
		new(productmodel.Product),
	)
}
