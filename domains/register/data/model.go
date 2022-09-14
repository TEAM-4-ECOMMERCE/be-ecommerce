package data

import (
	// "e-commerce/domains/users/entity"
	entity "e-commerce/domains/register/entity"

	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	Name		string
	Email		string
	Password	string
}

func FromCoreRegister(dataCore entity.Registers) Register {
	dataModel := Register {
		Name: dataCore.Name,
		Email: dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel
}

func (data *Register) ToCoreRegister() entity.Registers {
	return entity.Registers{
		UID: 		int(data.ID),
		Name: 		data.Name,
		Email: 		data.Email,
		Password: 	data.Password,
	}
}

func CoreListRegister(data []Register) []entity.Registers{
	var DataCore []entity.Registers
	for key := range data {
		DataCore = append(DataCore, data[key].ToCoreRegister())
	}

	return DataCore
}
