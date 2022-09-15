package transactionhandlers

import (
	entity "e-commerce/domains/transaction/entity"
)

type StoreRequest struct {
	Street      string  `json:"street"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zipcode     string  `json:"zipcode"`
	Visa        string  `json:"visa"`
	NameOnCard  string  `json:"name_on_card"`
	NumberCard  string  `json:"number_card"`
	ExpiredDate float64 `json:"expiration_date"`
}

type UpdateRequest struct {
	Status string
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
			NumberCard:  request.NameOnCard,
			ExpiredDate: request.ExpiredDate,
		},
	}
}
