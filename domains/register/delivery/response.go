package delivery

import users "e-commerce/domains/users/entity"

type ResponseRegister struct{
	Name		string
	Email		string
	Password	string
}

func FromCore(data users.Users) ResponseRegister {
	return ResponseRegister{
		Name: data.Name,
		Email: data.Email,
	}
}

func CoreList(data []users.Users) []ResponseRegister {
	var ResponseData []ResponseRegister
	for _, v := range data {
		ResponseData = append(ResponseData, FromCore(v))
	}
	return ResponseData
}

func CoreResponse(data users.Users) ResponseRegister {
	ResponseData := ResponseRegister{
		Name: data.Name,
		Email: data.Email,
	}
	return ResponseData
}