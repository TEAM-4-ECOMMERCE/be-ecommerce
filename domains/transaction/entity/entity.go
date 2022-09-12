package transactionentity

type TransactionEntity struct {
	TransactionID     uint
	TransactionCode   string
	GrandTotal        float64
	StatusTransaction string
	TransactionDate   float64
	Address           AddressEntity
	CreditCard        CreditCardEntity
	TransactionDetail []TransactionDetailEntity
}

type AddressEntity struct {
	AddressID uint
	Street    string
	City      string
	Province  string
	Zipcode   string
}

type CreditCardEntity struct {
	CreditCardID uint
	Visa         string
	NameOfCard   string
	NumberCard   string
	ExpiredDate  float64
}

type TransactionDetailEntity struct {
	TransactionDetailID uint
	Qty                 uint
	Subtotal            float64
	Product             ProductEntity
}

type ProductEntity struct {
	ProductID uint
	Name      string
	ImageUrl  string
}

type IusecaseTransaction interface {
	Store(transaction TransactionEntity) (err error)
	Update(transaction TransactionEntity) (err error)
	GetList(transaction TransactionEntity) (err error, result []TransactionDetailEntity)
	GetSingle(transaction TransactionEntity) (err error, result TransactionDetailEntity)
}

type IrepoTransaction interface {
	Insert(transaction TransactionEntity) (err error)
	Update(transaction TransactionEntity) (err error)
	FindAll(transaction TransactionEntity) (err error, result []TransactionEntity)
	Find(transaction TransactionEntity) (err error, result TransactionEntity)
}
