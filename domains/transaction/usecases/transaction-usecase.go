package transactionusecase

import (
	entity "e-commerce/domains/transaction/entity"
	"e-commerce/utils/helpers"
	"errors"
	"time"
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

	if len(carts) < 1 {
		return errors.New("nothing on this cart")
	}

	transaction.TransactionDetail = carts

	lastID, err := u.Usecase.FindLastInsertedId()

	if err != nil {
		return err
	}

	var grandTotal float64
	for _, cart := range carts {
		grandTotal += cart.Subtotal
	}

	transaction.TransactionCode = helpers.TFCode(uint(lastID))
	transaction.StatusTransaction = "on_delivery"
	transaction.GrandTotal = grandTotal
	transaction.TransactionDate = int64(time.Now().Unix())

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

	if err != nil {
		return result, err
	}

	for _, transaction := range transactionList {
		var grandTotal float64
		for _, detailTransaction := range transaction.TransactionDetail {
			grandTotal += detailTransaction.Subtotal
		}
		transaction.GrandTotal = grandTotal
	}

	return transactionList, nil
}

func (u *transactionUsecase) GetSingle(transaction entity.TransactionEntity) (result entity.TransactionEntity, err error) {
	return u.Usecase.Find(transaction)
}
