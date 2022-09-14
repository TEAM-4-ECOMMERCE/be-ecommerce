package carthandler

type cartRequestStore struct {
	ProductID uint `json:"product_id"`
}

type cartRequestUpdate struct {
	Qty uint `json:"qty"`
}
