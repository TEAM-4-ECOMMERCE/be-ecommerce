package producthandler

import (
	entity "e-commerce/domains/product/entity"
)

type ProductRequest struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
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
