package cartusecase

import (
	entity "e-commerce/domains/cart/entity"
	"e-commerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	entityRequest := entity.CartEntity{
		ProductID: 1,
	}

	entityResponse := entity.CartEntity{
		CartID:    2,
		ProductID: 2,
	}

	t.Run("Success store cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		entityResponse.CartID = 0
		repo.On("Find", mock.Anything).Return(entityResponse, nil).Once()
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Store(entityRequest)

		assert.NoError(t, err)
	})

	t.Run("Failed find data cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("Find", mock.Anything).Return(entity.CartEntity{}, errors.New("error find data")).Once()
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Store(entityRequest)

		assert.Error(t, err)
	})

	t.Run("Duplicate data", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		entityResponse.CartID = 1
		repo.On("Find", mock.Anything).Return(entityResponse, nil).Once()
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Store(entityRequest)

		assert.NoError(t, err)
	})

	t.Run("Insert Error", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		entityResponse.CartID = 0
		repo.On("Find", mock.Anything).Return(entityResponse, nil).Once()
		repo.On("Insert", mock.Anything).Return(0, errors.New("internal server error")).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Store(entityRequest)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	entityRequest := entity.CartEntity{
		CartID:    1,
		ProductID: 1,
		UserID:    1,
		Qty:       2,
	}

	t.Run("Success Update cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("Update", mock.Anything).Return(1, nil).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Update(entityRequest)

		assert.NoError(t, err)
	})

	t.Run("Error Update Cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("Update", mock.Anything).Return(0, errors.New("no updated")).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Update(entityRequest)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	entityRequest := entity.CartEntity{
		CartID:    1,
		ProductID: 1,
		UserID:    1,
		Qty:       2,
	}

	t.Run("Success Delete cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Delete(entityRequest)

		assert.NoError(t, err)
	})

	t.Run("Error Delete Cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("Delete", mock.Anything).Return(0, errors.New("no Deleted")).Once()

		cartUseCase := New(repo)
		err := cartUseCase.Delete(entityRequest)

		assert.Error(t, err)
	})
}

func TestGetList(t *testing.T) {
	entityRequest := entity.CartEntity{
		CartID:    1,
		ProductID: 1,
		UserID:    1,
		Qty:       2,
	}

	entityResponses := []entity.CartEntity{
		{
			ProductID: 2,
			Subtotal:  10000,
		},
		{
			ProductID: 3,
			Subtotal:  20000,
		},
	}

	t.Run("Success Get List cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("FindAll", mock.Anything).Return(entityResponses, nil).Once()

		cartUseCase := New(repo)
		carts, err := cartUseCase.GetList(entityRequest)

		assert.NoError(t, err)
		assert.Equal(t, float64(30000), carts[0].GrandTotal)
		assert.Equal(t, 2, len(carts))
	})

	t.Run("Error Get List Cart", func(t *testing.T) {
		repo := new(mocks.CartRepoMock)
		repo.On("FindAll", mock.Anything).Return([]entity.CartEntity{}, errors.New("internal server error")).Once()

		cartUseCase := New(repo)
		carts, err := cartUseCase.GetList(entityRequest)

		assert.Error(t, err)
		assert.Equal(t, 0, len(carts))
	})
}
