package Register

import "time"

type Registers struct {
	UID       int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IusecaseRegister interface {
	CreateUser(dataCreate Registers) (row int, err error)
}

type IregisterInterface interface {
	CreateUser(dataCreate Registers) (row int, err error)
}
