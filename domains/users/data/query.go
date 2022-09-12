package data

import (
	"e-commerce/domains/users/entity"
	"errors"

	"gorm.io/gorm"
)

type DataUser struct {
	db *gorm.DB
}

func NewDataBase(db *gorm.DB) users.UserInterface {
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

func (file *DataUser) DeleteData(dataDelete users.UserCore) (int, error)  {
	tx := file.db.Delete(&User{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0{
		return 0, errors.New("failed to delete user")
	}
	return int(tx.RowsAffected), nil
}

func (file *DataUser)UpdateData(dataUpdate users.UserCore) (int, error)  {
	tx := file.db.Model(&User{}).Updates(FromCore(dataUpdate))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return int(tx.RowsAffected), nil
}