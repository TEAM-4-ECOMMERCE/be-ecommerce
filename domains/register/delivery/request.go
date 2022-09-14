package delivery

import (
	register "e-commerce/domains/register/entity"

)

type requestRegister struct{
	Name		string 	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}


func fromCoreRequest(data requestRegister) register.Registers {
	return register.Registers{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}

func toCoreRequest(data requestRegister) register.Registers {
	return register.Registers{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}