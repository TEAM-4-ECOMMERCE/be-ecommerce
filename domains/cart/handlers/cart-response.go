package carthandler

import (
	entity "e-commerce/domains/cart/entity"
)

type cartResponse struct {
	CartID       uint    `json:"cart_id"`
	ProductName  string  `json:"product_name"`
	ImageUrl     string  `json:"image_url"`
	ProductPrice float64 `json:"product_price"`
	Qty          uint    `json:"qty"`
	Subtotal     float64 `json:"subtotal"`
}

func EntityToResponse(cartEntity entity.CartEntity) cartResponse {
	return cartResponse{
		CartID:       cartEntity.CartID,
		ProductName:  cartEntity.ProductName,
		ImageUrl:     cartEntity.ProductImageUrl,
		ProductPrice: cartEntity.ProductPrice,
		Qty:          cartEntity.Qty,
		Subtotal:     cartEntity.Subtotal,
	}
}
