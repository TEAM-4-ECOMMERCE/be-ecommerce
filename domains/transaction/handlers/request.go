package transactionhandlers

import (
	entity "e-commerce/domains/transaction/entity"
)

type StoreRequest struct {
	Street      string  `json:"street" validate:"required"`
	City        string  `json:"city" validate:"required"`
	State       string  `json:"state" validate:"required"`
	Zipcode     string  `json:"zipcode" validate:"required"`
	Visa        string  `json:"visa" validate:"required"`
	CVV         string  `json:"cvv" validate:"required"`
	NameOnCard  string  `json:"name_on_card" validate:"required"`
	NumberCard  string  `json:"number_card" validate:"required"`
	ExpiredDate float64 `json:"expired_date" validate:"required"`
}

type UpdateRequest struct {
	Status string `json:"status" validate:"required"`
}

func RequestToEntity(request StoreRequest) entity.TransactionEntity {
	return entity.TransactionEntity{
		Address: entity.AddressEntity{
			Street:   request.Street,
			City:     request.City,
			Province: request.State,
			Zipcode:  request.Zipcode,
		},
		CreditCard: entity.CreditCardEntity{
			Visa:        request.Visa,
			NameOfCard:  request.NameOnCard,
			NumberCard:  request.NumberCard,
			ExpiredDate: request.ExpiredDate,
			CVV:         request.CVV,
		},
	}
}
