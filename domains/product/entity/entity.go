package productentity

type ProductEntity struct {
	ProductID    uint
	UserID       uint
	CategoryID   uint
	CategoryName string
	Name         string
	Price        float64
	Stock        uint
	Description  string
	ImageUrl     string
	PreviousPage string
	CurrentPage  uint
	NextPage     string
	PageSize     uint
	SearchQuery  string
}

type IusecaseProduct interface {
	Store(product ProductEntity) (err error)
	Update(product ProductEntity) (err error)
	Delete(product ProductEntity) (err error)
	GetList(product ProductEntity) (result []ProductEntity, err error)
	GetSingle(product ProductEntity) (result ProductEntity, err error)
}

type IrepoProduct interface {
	Insert(product ProductEntity) (affectedRow int, err error)
	Update(product ProductEntity) (affectedRow int, err error)
	Delete(product ProductEntity) (affectedRow int, err error)
	CountData(uid uint) (totalRow uint, err error)
	FindAll(product ProductEntity) (result []ProductEntity, err error)
	Find(product ProductEntity) (result ProductEntity, err error)
}
