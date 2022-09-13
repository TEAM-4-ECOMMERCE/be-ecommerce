package Register

import "e-commerce/domains/users/entity"



type UserInterface interface {
	CreateUser (dataCreate users.UserCore) (row int, err error)
}