package cartmodel

import (
	entity "e-commerce/domains/cart/entity"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Qty       uint
	Subtotal  float64
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

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

func EntityToModel(entity entity.CartEntity) Cart {
	return Cart{
		ProductID: entity.ProductID,
		UserID:    entity.UserID,
		Qty:       entity.Qty,
		Subtotal:  entity.Subtotal,
	}
}

func ModelToEntity(model Cart) entity.CartEntity {
	return entity.CartEntity{
		CartID:          model.ID,
		ProductID:       model.ID,
		ProductName:     model.Product.Name,
		ProductPrice:    model.Product.Price,
		ProductImageUrl: model.Product.ImageUrl,
		Qty:             model.Qty,
		Subtotal:        float64(model.Qty) * model.Product.Price,
	}
}
