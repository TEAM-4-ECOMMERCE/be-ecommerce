package authhandler

import (
	entity "e-commerce/domains/auth/entity"
)

type authRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func requestToEntity(request authRequest) entity.AuthEntity {
	return entity.AuthEntity{
		Email:    request.Email,
		Password: request.Password,
	}
}
