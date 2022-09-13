package delivery

import users "e-commerce/domains/users/entity"

type UserRequest struct{
	Name		string 	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

func fromCoreRequest(data UserRequest) users.UserCore {
	return users.UserCore{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}

func toCoreRequest(data UserRequest) users.UserCore {
	return users.UserCore{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}