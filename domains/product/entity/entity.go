package productentity

type ProductEntity struct {
	ProductID    string
	UserID       uint
	Name         string
	Category     CategoryEntity
	Price        float64
	Stock        uint
	Description  string
	ImageUrl     string
	PreviousPage string
	CurrentPage  string
	NextPage     string
}

type CategoryEntity struct {
	CategoryID   int
	CategoryName string
}

type IusecaseProduct interface {
	Store(product ProductEntity) (err error)
	Update(product ProductEntity) (err error)
	Delete(product ProductEntity) (err error)
	GetList(product ProductEntity) (err error, result []ProductEntity)
	GetSingle(product ProductEntity) (err error, result ProductEntity)
}

type IrepoProduct interface {
	Insert(cart ProductEntity) (err error)
	Update(cart ProductEntity) (err error)
	Delete(cart ProductEntity) (err error)
	FindAll(cart ProductEntity) (err error, result []ProductEntity)
	Find(cart ProductEntity) (err error, result ProductEntity)
}
