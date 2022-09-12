package authentity

type AuthEntity struct {
	Email    string
	Password string
}

type IusecaseAuth interface {
	Login(userData AuthEntity) (err error, token string)
}

type IrepoAuth interface {
	SelectUserByEmail(email string) (err error, uid int)
}
