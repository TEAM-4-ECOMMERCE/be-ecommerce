package data

import (
	users "e-commerce/domains/users/entity"
	"errors"

	"gorm.io/gorm"
)

type DataUser struct {
	db *gorm.DB
}

func NewDataBase(db *gorm.DB) users.IuserInterface {
	return &DataUser{
		db: db,
	}
}

func (file *DataUser) SelectUser() ([]users.Users, error) {
	var data []User
	tx := file.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	Users := CoreList(data)
	return Users, nil
}

func (file *DataUser) DeleteUser(dataDelete users.Users) (int, error) {
	userModel := User{}
	userModel.ID = uint(dataDelete.UID)
	tx := file.db.Model(&User{}).Delete(&userModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete user")
	}
	return int(tx.RowsAffected), nil
}

func (file *DataUser) UpdateUser(dataUpdate users.Users) (int, error) {
	userModel := FromCore(dataUpdate)
	tx := file.db.Model(&User{}).Where("id = ?", dataUpdate.UID).Updates(&userModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return int(tx.RowsAffected), nil
}
