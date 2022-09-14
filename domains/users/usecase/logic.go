package usecase

import (
	users "e-commerce/domains/users/entity"
	"errors"
)

type useCase struct {
	userData users.IuserInterface
}

func NewLogic(data users.IuserInterface) users.IusecaseUser {
	return &useCase{
		userData: data,
	}
}

func (useCase *useCase) GetUser() ([]users.Users, error) {
	result, err := useCase.userData.SelectUser()
	return result, err
}

func (useCase *useCase) UpdateUser(dataUpdate users.Users) (int, error) {
	dataMap := make(map[string]interface{})
	if dataUpdate.Name != "" {
		dataMap["name"] = &dataUpdate.Name
	}
	if dataUpdate.Email != "" {
		dataMap["email"] = &dataUpdate.Email
	}
	if dataUpdate.ImageUrl != "" {
		dataMap["image_url"] = &dataUpdate.Email
	}

	result, err := useCase.userData.UpdateUser(dataUpdate)
	if err != nil {
		return 0, errors.New("user already register? you can't update if your account not listed")
	}
	return result, err
}

func (useCase *useCase) DeleteUser(dataDelete users.Users) (row int, err error) {
	result, err := useCase.userData.DeleteUser(dataDelete)
	if err != nil {
		return -1, errors.New("what you delete? register account first and then try again delete account")
	}
	return result, err
}
