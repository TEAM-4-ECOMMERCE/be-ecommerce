package transactionmodel

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID       uint
	AddressID    uint
	CreditCardID uint
	GrandTotal   float64
	Status       string

	TransactionDetail []TransactionDetail
	Address           Address
	CreditCard        CreditCard
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	ProductID     uint
	Qty           uint
	Subtotal      float64

	Product Product
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
	Zip    string
}

type CreditCard struct {
	gorm.Model
	Type   string
	Name   string
	Number string
	CVV    string
	Date   float64
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
