package data

import (
	"gorm.io/gorm"
	category "e-commerce/domains/category/entity"
)
type Category struct {
	gorm.Model
	Name string
}

func FromCore(data category.CategoryEntity) Category {
	dataModel := Category{
		Name: 	data.CategoryName,
	}
	return dataModel
}


func (data *Category) toCore() category.CategoryEntity{
	return category.CategoryEntity{
		CategoryID: 	int(data.ID),
		CategoryName:	data.Name,
	}
}

func CoreList(data []Category) []category.CategoryEntity {
	var CategoryCore []category.CategoryEntity
	for key := range data {
		CategoryCore = append(CategoryCore, data[key].toCore())
	}
	return CategoryCore
}