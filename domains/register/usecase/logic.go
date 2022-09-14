package usecase

import (
	reg "e-commerce/domains/register/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type useCase struct {
	userData reg.IusecaseRegister
}

func NewLogic(data reg.IusecaseRegister) reg.IusecaseRegister {
	return &useCase{
		userData: data,
	}
}

func (useCase *useCase) CreateUser(dataCreate reg.Registers) (int, error) {
	if dataCreate.Name == "" || dataCreate.Email == "" || dataCreate.Password == "" {
		return -1, errors.New("all data must be filled")
	}

	passWillBcrypt := []byte(dataCreate.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return -2, errors.New("hashing password failed")
	}
	dataCreate.Password = string(hash)
	result, err := useCase.userData.CreateUser(dataCreate)
	if err != nil {
		return 0, errors.New("failed to insert data")
	}
	return result, nil
}