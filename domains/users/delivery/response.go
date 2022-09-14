package delivery

import users "e-commerce/domains/users/entity"

type UserResponse struct {
	Name		string
	Email		string
}

func FromCore(data users.Users) UserResponse {
	return UserResponse{
		Name: data.Name,
		Email: data.Email,
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
		Name: data.Name,
		Email: data.Email,
	}
	return ResponseData
}