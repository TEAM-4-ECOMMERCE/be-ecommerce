package utils

import (
	"e-commerce/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	cartModel "e-commerce/domains/cart/models"
	categoryModel "e-commerce/domains/category/model"
	productmodel "e-commerce/domains/product/models"
	users "e-commerce/domains/users/data"
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
		new(users.User),
		new(productmodel.Product),
		new(cartModel.Cart),
	)
}
