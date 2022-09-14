package delivery

import users "e-commerce/domains/users/entity"

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

func FromCore(data users.Users) UserResponse {
	return UserResponse{
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
	}
}

func CoreList(data []users.Users) []UserResponse {
	var ResponseData []UserResponse
	for _, v := range data {
		ResponseData = append(ResponseData, FromCore(v))
	}
	return ResponseData
}

func CoreResponse(data users.Users) UserResponse {
	ResponseData := UserResponse{
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
	}
	return ResponseData
}
