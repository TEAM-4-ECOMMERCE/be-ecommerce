package producthandler

import (
	entity "e-commerce/domains/product/entity"
)

type ProductRequest struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
}

func RequestToEntity(request ProductRequest) entity.ProductEntity {
	return entity.ProductEntity{
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		Price:       float64(request.Price),
		Stock:       request.Stock,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
	}
}
