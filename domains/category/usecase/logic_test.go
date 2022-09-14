package usecase

import (
	category "e-commerce/domains/category/entity"
	"e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetCategory(t *testing.T)()  {
	repo := new(mocks.CategoryRepository)
	returnData := []category.CategoryEntity{{CategoryID: 1, CategoryName: "Ikebal"}}

	t.Run("Succses Get Category", func(t *testing.T)  {
		repo.On("GetCategory").Return(returnData, nil)

		usecase := New(repo)
		resultData, err := usecase.GetCategory()
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].CategoryID, returnData[0].CategoryID)
		repo.AssertExpectations(t)
	})
}