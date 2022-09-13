package data

import (
	register "e-commerce/domains/register/entity"
	"e-commerce/domains/users/data"
	"e-commerce/domains/users/entity"

	"gorm.io/gorm"
)

type Receive struct{
	db *gorm.DB
}

func NewRegister(db *gorm.DB) register.UserInterface {
	return &Receive {
		db: db,
	}
}

func (file *Receive) CreateUser(dataCreate users.UserCore) (int, error) {
	UserModel := data.FromCore(dataCreate)
	tx := file.db.Create(&UserModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}