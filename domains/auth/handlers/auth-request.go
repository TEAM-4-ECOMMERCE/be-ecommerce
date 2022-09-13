package authhandler

import (
	entity "e-commerce/domains/auth/entity"
)

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func requestToEntity(request authRequest) entity.AuthEntity {
	return entity.AuthEntity{
		Email:    request.Email,
		Password: request.Password,
	}
}
