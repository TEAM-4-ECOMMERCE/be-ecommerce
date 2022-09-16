package carthandler

type cartRequestStore struct {
	ProductID uint `json:"product_id" validate:"required"`
}

type cartRequestUpdate struct {
	Qty uint `json:"qty" validate:"required"`
}
