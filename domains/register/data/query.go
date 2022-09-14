package data

import (
	reg "e-commerce/domains/register/entity"
	"e-commerce/middlewares"
	"fmt"
	"strconv"

	// "errors"

	"gorm.io/gorm"
)

type Receive struct {
	db *gorm.DB
}

func NewRegister(db *gorm.DB) reg.IregisterInterface {
	return &Receive{
		db: db,
	}
}

func (file *Receive) CreateUser(dataCreate reg.Registers) (int, error) {
	UserModel := FromCoreRegister(dataCreate)
	tx := file.db.Create(&UserModel)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return 0, tx.Error
	}
	token, errToken := middlewares.CreateToken(UserModel.ID)
	tokenConv, _ := strconv.Atoi(token)
	if errToken != nil {
		return -1, errToken
	}
	return tokenConv, nil
}
