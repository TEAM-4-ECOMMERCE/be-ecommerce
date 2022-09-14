package cartentity

type CartEntity struct {
	CartID            uint
	UserID            uint
	Qty               uint
	Subtotal          float64
	Product           ProductEntity
	TotalOrderProduct uint
	GrandTotal        float64	
}

type ProductEntity struct {
	ProductID  uint
	CategoryID uint
	UserID     uint
	Name       string
	Price      float64
	Stock      uint
	Desc       string
	Image      string
}

type IusecaseProduct interface {
	Store(cart CartEntity) (err error)
	Update(cart CartEntity) (err error)
	Delete(cart CartEntity) (err error)
	GetList(cart CartEntity) (err error, result []CartEntity)
	GetSingle(cart CartEntity) (err error, result CartEntity)
}

type IrepoProduct interface {
	Insert(cart CartEntity) (err error)
	Update(cart CartEntity) (err error)
	Delete(cart CartEntity) (err error)
	FindAll(cart CartEntity) (err error, result []CartEntity)
	Find(cart CartEntity) (err error, result CartEntity)
}
