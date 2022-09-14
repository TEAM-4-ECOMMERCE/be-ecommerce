package usecase

import (
	users "e-commerce/domains/users/entity"
	"e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetData(t *testing.T)()  {
	repo := new(mocks.UserRepository)
	returnData := []users.Users{{UID: 1, Name: "Iqbal",Email: "iqbalinaja@yahoo.com",Password: "abogoboga"}}

	t.Run("Succses Get Data", func(t *testing.T)  {
		repo.On("SelectUser").Return(returnData, nil)

		usecase := NewLogic(repo)
		resultData, err := usecase.SelectUser()
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].UID, returnData[0].UID)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T)()  {
	repo := new(mocks.UserRepository)
	t.Run("Test Update Data", func(t *testing.T) {
		repo.On("UpdateUser", mock.Anything).Return(1, nil).Once()
		data := users.Users{
			Name:     "Ikebal",
			Email:    "maklogeming@gmail.com",
			Password: "qwerty",
		}
		usercase := NewLogic(repo)
		result, err := usercase.UpdateUser(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T)()  {
	repo := new(mocks.UserRepository)
	t.Run("Test Delete Data", func(t *testing.T) {
		repo.On("DeleteUser", mock.Anything).Return(1, nil).Once()
		data := users.Users{
			Name:     "Ikebal",
		}
		usercase := NewLogic(repo)
		result, err := usercase.DeleteUser(data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
}