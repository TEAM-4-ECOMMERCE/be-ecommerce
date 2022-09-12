package data

import (
	"e-commerce/domains/users/entity"

	"gorm.io/gorm"
)

type DataUser struct {
	db *gorm.DB
}

func DataBase(db *gorm.DB) users.UserInterface {
	return &DataUser {
		db: db,
	}
}

func (file *DataUser)GetUser() ([]users.UserCore, error)  {
	var data []User
	tx := file.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	UserCore := CoreList(data)
	return UserCore, nil
}

func (file *DataUser) CreateUser(dataCreate users.UserCore) (int, error) {
	UserModel := FromCore(dataCreate)
	tx := file.db.Create(&UserModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}