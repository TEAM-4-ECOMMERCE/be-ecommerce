package transactionhandlers

import (
	entity "e-commerce/domains/transaction/entity"
)

type ListResponse struct {
	TransactionID     uint    `json:"transaction_id"`
	UserID            uint    `json:"user_id"`
	TransactionCode   string  `json:"transaction_code"`
	GrandTotal        float64 `json:"grand_total"`
	StatusTransaction string  `json:"status_order"`
	TransactionDate   int64   `json:"transaction_date"`
}

type SingleResponse struct {
	TransactionID     uint                        `json:"transaction_id"`
	UserID            uint                        `json:"user_id"`
	TransactionCode   string                      `json:"transaction_code"`
	GrandTotal        float64                     `json:"grand_total"`
	StatusTransaction string                      `json:"status"`
	TransactionDate   int64                       `json:"transaction_date"`
	Address           AddressResponse             `json:"address"`
	CreditCard        CreditCardResponse          `json:"credit_card"`
	TransactionDetail []TransactionDetailResponse `json:"products"`
}

type AddressResponse struct {
	AddressID uint   `json:"address_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Zipcode   string `json:"zipcode"`
}

type CreditCardResponse struct {
	CreditCardID uint    `json:"credit_card_id"`
	Visa         string  `json:"visa"`
	NameOfCard   string  `json:"name_of_card"`
	NumberCard   string  `json:"number_card"`
	CVV          string  `json:"cvv"`
	ExpiredDate  float64 `json:"expired_date"`
}

type TransactionDetailResponse struct {
	TransactionDetailID uint    `json:"transaction_detail_id"`
	Qty                 uint    `json:"qty"`
	Subtotal            float64 `json:"subtotal"`
	ProductID           uint    `json:"product_id"`
	ProductName         string  `json:"product_name"`
	ProductPrice        float64 `json:"product_price"`
	ImageUrl            string  `json:"image_url"`
}

func EntityToResponseSingle(Entity entity.TransactionEntity) SingleResponse {
	var ListDetailTransaction []TransactionDetailResponse

	for _, DetailTransaction := range Entity.TransactionDetail {
		ListDetailTransaction = append(ListDetailTransaction, TransactionDetailResponse{
			TransactionDetailID: DetailTransaction.TransactionDetailID,
			Qty:                 DetailTransaction.Qty,
			Subtotal:            DetailTransaction.Subtotal,
			ProductID:           DetailTransaction.ProductID,
			ProductName:         DetailTransaction.ProductName,
			ProductPrice:        DetailTransaction.ProductPrice,
			ImageUrl:            DetailTransaction.ImageUrl,
		})
	}

	return SingleResponse{
		TransactionID:     Entity.TransactionID,
		UserID:            Entity.UserID,
		TransactionCode:   Entity.TransactionCode,
		GrandTotal:        Entity.GrandTotal,
		StatusTransaction: Entity.StatusTransaction,
		TransactionDate:   Entity.TransactionDate,
		Address: AddressResponse{
			AddressID: Entity.Address.AddressID,
			Street:    Entity.Address.Street,
			City:      Entity.Address.City,
			Province:  Entity.Address.Province,
			Zipcode:   Entity.Address.Zipcode,
		},
		CreditCard: CreditCardResponse{
			CreditCardID: Entity.CreditCard.CreditCardID,
			Visa:         Entity.CreditCard.Visa,
			NumberCard:   Entity.CreditCard.NumberCard,
			NameOfCard:   Entity.CreditCard.NameOfCard,
			ExpiredDate:  Entity.CreditCard.ExpiredDate,
			CVV:          Entity.CreditCard.CVV,
		},
		TransactionDetail: ListDetailTransaction,
	}
}

func EntityToResponseList(Entity []entity.TransactionEntity) []ListResponse {
	var listTransaction []ListResponse

	for _, transaction := range Entity {
		listTransaction = append(listTransaction, ListResponse{
			TransactionID:     transaction.TransactionID,
			UserID:            transaction.UserID,
			TransactionCode:   transaction.TransactionCode,
			GrandTotal:        transaction.GrandTotal,
			StatusTransaction: transaction.StatusTransaction,
			TransactionDate:   transaction.TransactionDate,
		})
	}

	return listTransaction
}
