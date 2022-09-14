package delivery

import category"e-commerce/domains/category/entity"

type CategoryResponse struct {
	CategoryName	string
}

func FromCore(data category.CategoryEntity) CategoryResponse {
	return CategoryResponse{
		CategoryName: 	data.CategoryName,
	}
}

func CoreList(data []category.CategoryEntity) []CategoryResponse {
	var Response []CategoryResponse
	for _, v := range data {
		Response = append(Response, FromCore(v))
	}
	return Response
}

func CoreResponse(data category.CategoryEntity) CategoryResponse {
	Response := CategoryResponse{
		CategoryName: data.CategoryName,
	}
	return Response
}