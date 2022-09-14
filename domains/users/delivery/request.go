package delivery

import users "e-commerce/domains/users/entity"

type UserRequest struct {
	UserID   int    `json:"id" form:"id"`
	Name     string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	ImageUrl string `json:"image_url" from:"image_url"`
}

func FromCoreRequest(data UserRequest) users.Users {
	return users.Users{
		UID:      data.UserID,
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
	}
}

func toCoreRequest(data UserRequest) users.Users {
	return users.Users{
		UID:      data.UserID,
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.ImageUrl,
	}
}
