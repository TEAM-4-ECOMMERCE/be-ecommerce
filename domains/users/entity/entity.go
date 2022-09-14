package users

import "time"

type Users struct {
	UID       int
	Name      string
	Email     string
	Password  string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IusecaseUser interface {
	GetUser() (data []Users, err error)
	DeleteUser(dataDelete Users) (row int, err error)
	UpdateUser(dataUpdate Users) (row int, err error)
}

type IuserInterface interface {
	SelectUser() (data []Users, err error)
	DeleteUser(dataDelete Users) (row int, err error)
	UpdateUser(dataUpdate Users) (row int, err error)
}
