package delivery

import category "e-commerce/domains/category/entity"

type CategoryRequest struct {
	CategoryName	string
}

func fromCore(data CategoryRequest) category.CategoryEntity {
	return category.CategoryEntity{
		CategoryName: data.CategoryName,
	}
}