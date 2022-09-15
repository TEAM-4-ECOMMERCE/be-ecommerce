package transactionusecase

import (
	entity "e-commerce/domains/transaction/entity"
	"e-commerce/utils/helpers"
)

type transactionUsecase struct {
	Usecase entity.IrepoTransaction
}

func New(usecase entity.IrepoTransaction) *transactionUsecase {
	return &transactionUsecase{
		Usecase: usecase,
	}
}

func (u *transactionUsecase) Store(transaction entity.TransactionEntity) (err error) {

	carts, err := u.Usecase.FindCart(transaction)

	if err != nil {
		return err
	}

	transaction.TransactionDetail = carts

	lastID, err := u.Usecase.FindLastInsertedId()

	if err != nil {
		return err
	}

	transaction.TransactionCode = helpers.TFCode(uint(lastID))

	err = u.Usecase.Insert(transaction)

	if err != nil {
		return err
	}

	err = u.Usecase.DeleteCart(transaction)

	if err != nil {
		return err
	}

	return nil
}

func (u *transactionUsecase) Update(transaction entity.TransactionEntity) (err error) {
	err = u.Usecase.Update(transaction)

	if err != nil {
		return err
	}

	return nil
}

func (u *transactionUsecase) GetList(transaction entity.TransactionEntity) (result []entity.TransactionEntity, err error) {
	transactionList, err := u.Usecase.FindAll(transaction)

	for _, transaction := range transactionList {
		var grandTotal float64
		for _, detailTransaction := range transaction.TransactionDetail {
			grandTotal += detailTransaction.Subtotal
		}
		transaction.GrandTotal = grandTotal
	}

	return
}

func (u *transactionUsecase) GetSingle(transaction entity.TransactionEntity) (result entity.TransactionEntity, err error) {
	result, err = u.Usecase.Find(transaction)
	if err != nil {
		return result, err
	}

	var grandTotal float64
	for _, transactionDetail := range result.TransactionDetail {
		grandTotal += transactionDetail.Subtotal
	}

	result.GrandTotal = grandTotal
	return
}
