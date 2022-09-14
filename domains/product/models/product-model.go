package productmodel

import (
	entity "e-commerce/domains/product/entity"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  uint
	UserID      uint
	Name        string
	Price       float64
	Stock       uint
	Description string
	ImageUrl    string
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User        User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Category struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func ProductModelToEntity(productModel Product) entity.ProductEntity {
	return entity.ProductEntity{
		ProductID:    productModel.ID,
		UserID:       productModel.UserID,
		Name:         productModel.Name,
		Description:  productModel.Description,
		ImageUrl:     productModel.ImageUrl,
		Price:        productModel.Price,
		Stock:        productModel.Stock,
		CategoryID:   productModel.CategoryID,
		CategoryName: productModel.Category.Name,
	}
}

func ProductEntityToModel(productEntity entity.ProductEntity) Product {
	return Product{
		UserID:      productEntity.UserID,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		ImageUrl:    productEntity.ImageUrl,
		Price:       productEntity.Price,
		Stock:       productEntity.Stock,
		CategoryID:  productEntity.CategoryID,
	}
}
