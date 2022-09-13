package producthandler

import (
	entity "e-commerce/domains/product/entity"
)

type ProductResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	CategoryName string  `json:"category_name"`
	Price        float64 `json:"price"`
	Stock        uint    `json:"stock"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"image_url"`
}

func EntityToProductResponse(productEntity entity.ProductEntity) ProductResponse {
	return ProductResponse{
		ID:           productEntity.ProductID,
		Name:         productEntity.Name,
		CategoryName: productEntity.CategoryName,
		Price:        productEntity.Price,
		Stock:        productEntity.Stock,
		Description:  productEntity.Description,
		ImageUrl:     productEntity.ImageUrl,
	}
}
