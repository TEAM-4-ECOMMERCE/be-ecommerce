package authrepo

import (
	entity "e-commerce/domains/auth/entity"
	authUserModel "e-commerce/domains/auth/models"

	"gorm.io/gorm"
)

type authRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *authRepo {
	return &authRepo{
		DB: db,
	}
}

func (r *authRepo) SelectUserByEmail(email string) (authEntity entity.AuthEntity, err error) {
	userModel := authUserModel.User{}

	tx := r.DB.Where("email = ?", email).First(&userModel)
	err = tx.Error

	if err != nil {
		return authEntity, err
	}

	authEntity = authUserModel.ModelToEntity(userModel)

	return authEntity, err
}
