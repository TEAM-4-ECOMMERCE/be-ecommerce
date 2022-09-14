package data

import (
	// "e-commerce/domains/users/entity"
	entity "e-commerce/domains/register/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	ImageUrl string

	Cart []Cart
}

type Cart struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Qty       uint
	Subtotal  float64
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func FromCoreRegister(dataCore entity.Registers) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		ImageUrl: dataCore.ImageUrl,
	}
	return dataModel
}

func (data *User) ToCoreRegister() entity.Registers {
	return entity.Registers{
		UID:      int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		ImageUrl: data.ImageUrl,
	}
}

func CoreListRegister(data []User) []entity.Registers {
	var DataCore []entity.Registers
	for key := range data {
		DataCore = append(DataCore, data[key].ToCoreRegister())
	}

	return DataCore
}
