package usecase

import (
	users "e-commerce/domains/users/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type useCase struct {
	userData users.IuserInterface
}

func NewLogic(data users.IuserInterface) users.IuserInterface {
	return &useCase{
		userData: data,
	}
}

func (useCase *useCase) GetUser() ([]users.Users, error) {
	result, err := useCase.userData.GetUser()
	return result, err
}

func (useCase *useCase) UpdateUser(dataUpdate users.Users) (int, error) {
	dataMap := make(map[string]interface{})
	if dataUpdate.Name != ""{
		dataMap["name"] = &dataUpdate.Name
	}
	if dataUpdate.Email != ""{
		dataMap["email"] = &dataUpdate.Email
	}
	if dataUpdate.Password != ""{
		bytes, err := bcrypt.GenerateFromPassword([]byte(dataUpdate.Password), bcrypt.DefaultCost)
		if err != nil {
			return -1, errors.New("failed bcrypt the password")
		}
		dataMap["password"] = &bytes
	}


	result, err := useCase.userData.UpdateUser(dataUpdate)
	if err != nil{
		return 0, errors.New("user already register? you can't update if your account not listed")
	}
	return result, err
}

func (useCase *useCase) DeleteUser(dataDelete users.Users) (row int, err error) {
	result, err := useCase.userData.DeleteUser(dataDelete)
	if err != nil {
		return -1, errors.New("what you delete? register account first and then try again delete account")
	}
	return result , err 
}

