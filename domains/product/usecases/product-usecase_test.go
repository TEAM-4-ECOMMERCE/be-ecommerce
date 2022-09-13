package productusecase

import (
	entity "e-commerce/domains/product/entity"
	"e-commerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetList(t *testing.T) {
	repo := new(mocks.ProductRepoMock)

	productResponseEntityList := []entity.ProductEntity{
		{
			ProductID:    1,
			UserID:       1,
			CategoryID:   1,
			CategoryName: "Laptop",
			Name:         "Thinkpad X270",
			Price:        25000000,
			Stock:        2,
			Description:  "Awet Dan Bagus",
			ImageUrl:     "https://i0.wp.com/dikisepterian.com/wp-content/uploads/2022/02/ThinkPad-X270-diki-septerian-21.jpg?resize=640%2C410&ssl=1",
		},
		{
			ProductID:    2,
			UserID:       2,
			CategoryID:   2,
			CategoryName: "Laptop",
			Name:         "ASUS ROG",
			Price:        25000000,
			Stock:        5,
			Description:  "Awet Dan Bagus",
			ImageUrl:     "https://i0.wp.com/dikisepterian.com/wp-content/uploads/2022/02/ThinkPad-X270-diki-septerian-21.jpg?resize=640%2C410&ssl=1",
		},
	}

	produtRequestEntity := entity.ProductEntity{
		CurrentPage: 1,
		PageSize:    5,
		CategoryID:  1,
		SearchQuery: "aka",
	}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("FindAll", mock.Anything).Return(productResponseEntityList, nil).Once()

		productUsecase := New(repo)

		products, err := productUsecase.GetList(produtRequestEntity)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(products))

		repo.AssertExpectations(t)
	})

	t.Run("Current Page is 0", func(t *testing.T) {
		repo.On("FindAll", mock.Anything).Return(productResponseEntityList, nil).Once()

		productUsecase := New(repo)

		produtRequestEntity.CurrentPage = 0
		products, err := productUsecase.GetList(produtRequestEntity)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(products))

		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("FindAll", mock.Anything).Return([]entity.ProductEntity{}, errors.New("internal server error")).Once()

		productUsecase := New(repo)

		produtRequestEntity.CurrentPage = 0
		products, err := productUsecase.GetList(produtRequestEntity)

		assert.Error(t, err)
		assert.Equal(t, 0, len(products))

		repo.AssertExpectations(t)
	})
}

func TestGetSingle(t *testing.T) {
	repo := new(mocks.ProductRepoMock)

	productResponseEntity := entity.ProductEntity{
		ProductID:    1,
		UserID:       1,
		CategoryID:   1,
		CategoryName: "Laptop",
		Name:         "Thinkpad X270",
		Price:        25000000,
		Stock:        2,
		Description:  "Awet Dan Bagus",
		ImageUrl:     "https://i0.wp.com/dikisepterian.com/wp-content/uploads/2022/02/ThinkPad-X270-diki-septerian-21.jpg?resize=640%2C410&ssl=1",
	}

	productRequest := entity.ProductEntity{
		ProductID: 1,
	}

	t.Run("Success Get Single Data", func(t *testing.T) {
		repo.On("Find", mock.Anything).Return(productResponseEntity, nil).Once()

		productUsecase := New(repo)

		product, err := productUsecase.GetSingle(productRequest)

		assert.NoError(t, err)
		assert.Equal(t, productResponseEntity.UserID, product.UserID)
		repo.AssertExpectations(t)
	})

	t.Run("error get single data", func(t *testing.T) {
		repo.On("Find", mock.Anything).Return(entity.ProductEntity{}, errors.New("internal server error")).Once()

		productUsecase := New(repo)

		product, err := productUsecase.GetSingle(productRequest)

		assert.Error(t, err)
		assert.NotEqual(t, productResponseEntity.UserID, product.UserID)
		repo.AssertExpectations(t)
	})
}
