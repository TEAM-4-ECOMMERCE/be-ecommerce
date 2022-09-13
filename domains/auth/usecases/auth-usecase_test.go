package authusecase

import (
	authentity "e-commerce/domains/auth/entity"
	"e-commerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthentication(t *testing.T) {
	repo := new(mocks.AuthRepoMock)

	authEntityRequest := authentity.AuthEntity{
		Email:    "havis@gmail.com",
		Password: "shandrinaUwU",
	}

	authEntityResponse := authentity.AuthEntity{
		Id:       1,
		Email:    "havis@gmail.com",
		Password: "$2a$12$vRWUCiuuQhoy5iOXOk7FlO1/5c9rzHQHQ11YlMapsUrMQcO.DwTWy",
	}

	t.Run("Success Login", func(t *testing.T) {
		repo.On("SelectUserByEmail", mock.Anything).Return(authEntityResponse, nil).Once()

		usecaseImpl := New(repo)

		token, err := usecaseImpl.Login(authEntityRequest)

		assert.NoError(t, err)
		assert.NotEqual(t, "", token)
	})

	t.Run("Failed Login", func(t *testing.T) {
		repo.On("SelectUserByEmail", mock.Anything).Return(authentity.AuthEntity{}, errors.New("internal server error")).Once()

		usecaseImpl := New(repo)

		token, err := usecaseImpl.Login(authEntityRequest)

		assert.Error(t, err)
		assert.Equal(t, "", token)
	})
}
