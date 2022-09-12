package users

import "time"

type UserCore struct {
	UID			int
	Name		string
	Email		string
	Password	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type UserInterface interface {
	GetUser() (data []UserCore, err error)
	CreateUser (dataCreate UserCore) (row int, err error)
	// DeleteData (dataDelete UserCore) (row int, err error)
	// UpdateData (dataUpdate UserCore) (row int, err error)
}

