package authusecase

import (
	entity "e-commerce/domains/auth/entity"
	"e-commerce/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	repo entity.IrepoAuth
}

func New(repo entity.IrepoAuth) *authUseCase {
	return &authUseCase{
		repo: repo,
	}
}

func (u *authUseCase) Login(userData entity.AuthEntity) (token string, err error) {
	result, err := u.repo.SelectUserByEmail(userData.Email)

	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userData.Password))

	if err != nil {
		return token, err
	}

	return middlewares.CreateToken(result.Id)
}
