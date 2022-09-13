package authentity

type AuthEntity struct {
	Id       uint
	Email    string
	Password string
}

type IusecaseAuth interface {
	Login(userData AuthEntity) (token string, err error)
}

type IrepoAuth interface {
	SelectUserByEmail(email string) (userEntity AuthEntity, err error)
}
