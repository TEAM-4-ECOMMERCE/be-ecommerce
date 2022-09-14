package data

import (
	users "e-commerce/domains/users/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string

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

func FromCore(dataCore users.Users) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel
}

func (data *User) ToCore() users.Users {
	return users.Users{
		UID:      int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreList(data []User) []users.Users {
	var DataCore []users.Users
	for key := range data {
		DataCore = append(DataCore, data[key].ToCore())
	}

	return DataCore
}
