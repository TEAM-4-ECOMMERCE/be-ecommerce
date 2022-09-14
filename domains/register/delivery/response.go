package delivery

import users "e-commerce/domains/users/entity"

type ResponseRegister struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

func FromCore(data users.Users) ResponseRegister {
	return ResponseRegister{
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
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
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
	}
	return ResponseData
}
