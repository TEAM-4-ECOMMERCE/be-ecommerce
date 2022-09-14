package usecase

import(
	category "e-commerce/domains/category/entity"
)

type useCase struct{
	categoryData category.IrepoCategory
}

func New(data category.IrepoCategory) category.IrepoCategory {
	return &useCase{
		categoryData: data,
	}
}

func (useCase *useCase) GetCategory()([]category.CategoryEntity, error) {
	result, err := useCase.categoryData.GetCategory()
	return result, err
}