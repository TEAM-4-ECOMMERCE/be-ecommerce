package transactionentity

type TransactionEntity struct {
	TransactionID     uint
	UserID            uint
	TransactionCode   string
	GrandTotal        float64
	StatusTransaction string
	TransactionDate   int64
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
	CVV          string
	ExpiredDate  float64
}

type TransactionDetailEntity struct {
	TransactionDetailID uint
	Qty                 uint
	Subtotal            float64
	ProductID           uint
	ProductName         string
	ProductPrice        float64
	ImageUrl            string
}

type IusecaseTransaction interface {
	Store(transaction TransactionEntity) (err error)
	Update(transaction TransactionEntity) (err error)
	GetList(transaction TransactionEntity) (result []TransactionEntity, err error)
	GetSingle(transaction TransactionEntity) (result TransactionEntity, err error)
}

type IrepoTransaction interface {
	Insert(transaction TransactionEntity) (err error)
	Update(transaction TransactionEntity) (err error)
	DeleteCart(transaction TransactionEntity) (err error)
	FindAll(transaction TransactionEntity) (result []TransactionEntity, err error)
	Find(transaction TransactionEntity) (result TransactionEntity, err error)
	FindCart(transaction TransactionEntity) (result []TransactionDetailEntity, err error)
	FindLastInsertedId() (row int, err error)
}
