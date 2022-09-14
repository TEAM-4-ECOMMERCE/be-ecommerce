package utils

import (
	"e-commerce/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	categorymodel "e-commerce/domains/category/model"
	users "e-commerce/domains/users/data"
	reg "e-commerce/domains/register/data"
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
	db.AutoMigrate(new(categorymodel.Category))
	db.AutoMigrate(new(users.User))
	db.AutoMigrate(new(reg.Register))
}
