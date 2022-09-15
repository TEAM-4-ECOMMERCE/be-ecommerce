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
