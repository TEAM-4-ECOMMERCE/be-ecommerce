package delivery

import (
	register "e-commerce/domains/register/entity"
)

type requestRegister struct {
	Name     string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	ImageUrl string `json:"image_url" from:"image_url" validate:"required"`
}

func FromCoreRequest(data requestRegister) register.Registers {
	return register.Registers{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		ImageUrl: data.ImageUrl,
	}
}

func toCoreRequest(data requestRegister) register.Registers {
	return register.Registers{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		ImageUrl: data.ImageUrl,
	}
}
