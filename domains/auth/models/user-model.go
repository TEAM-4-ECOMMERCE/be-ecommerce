package authmodel

import (
	entity "e-commerce/domains/auth/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func ModelToEntity(userModel User) entity.AuthEntity {
	return entity.AuthEntity{
		Id:       userModel.ID,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}
