package data

import (
	"e-commerce/domains/users/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name		string
	Email		string
	Password	string
	// Cart		[]Carts
}



func FromCore(dataCore users.UserCore) User {
	dataModel := User{
		Name: dataCore.Name,
		Email: dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel
}

func (data *User) ToCore () users.UserCore  {
	return users.UserCore{
		UID: 		int(data.ID),
		Name: 		data.Name,
		Email: 		data.Email,
		Password: 	data.Password,
	}
}

func CoreList(data []User) []users.UserCore {
	var DataCore []users.UserCore
	for key := range data {
		DataCore = append(DataCore, data[key].ToCore())
	}

	return DataCore
}

