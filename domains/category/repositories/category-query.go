package data

import (
	category "e-commerce/domains/category/entity"
	datas "e-commerce/domains/category/data"

	"gorm.io/gorm"
)

type DataCategory struct {
	db *gorm.DB
}

func New(db *gorm.DB) category.IrepoCategory {
	return &DataCategory{
		db: db,
	}
}

func (file *DataCategory) GetCategory() ([]category.CategoryEntity, error) {
	var data []datas.Category
	tx := file.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	categorys := datas.CoreList(data)
	return categorys, nil
}
