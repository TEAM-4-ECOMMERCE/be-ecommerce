package cartentity

type CartEntity struct {
	CartID            uint
	ProductID         uint
	UserID            uint
	ProductName       string
	ProductImageUrl   string
	ProductPrice      float64
	Qty               uint
	Subtotal          float64
	TotalOrderProduct uint
	GrandTotal        float64
}

type IusecaseCart interface {
	Store(cart CartEntity) (err error)
	Update(cart CartEntity) (err error)
	Delete(cart CartEntity) (err error)
	GetList(cart CartEntity) (result []CartEntity, err error)
}

type IrepoCart interface {
	Insert(cart CartEntity) (affectedRow int, err error)
	Update(cart CartEntity) (affectedRow int, err error)
	Delete(cart CartEntity) (affectedRow int, err error)
	FindAll(cart CartEntity) (result []CartEntity, err error)
	Find(cart CartEntity) (result CartEntity, err error)
	FindProduct(cart CartEntity) (result CartEntity, err error)
}
