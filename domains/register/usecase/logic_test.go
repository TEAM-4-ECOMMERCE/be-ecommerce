package usecase

import (
	reg "e-commerce/domains/register/entity"
	"e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestInsertData(t *testing.T)()  {
	repo := new(mocks.RegisterRepository)
	t.Run("Test Create data succses", func(t *testing.T) {
		repo.On("CreateUser", mock.Anything).Return(1, nil).Once()
		data := reg.Registers{
			Name:     "Ikebal",
			Email:    "maklogeming@gmail.com",
			Password: "asdfghjkl",
		}
		usercase := NewLogic(repo)
		result, err := usercase.CreateUser(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
}