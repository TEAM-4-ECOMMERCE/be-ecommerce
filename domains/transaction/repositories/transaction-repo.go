package transactionrepositories

import (
	entity "e-commerce/domains/transaction/entity"
	model "e-commerce/domains/transaction/models"

	"gorm.io/gorm"
)

type transactionRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *transactionRepo {
	return &transactionRepo{
		DB: db,
	}
}

func (r *transactionRepo) Insert(transaction entity.TransactionEntity) (err error) {
	addressModel := model.Address{
		Street: transaction.Address.Street,
		City:   transaction.Address.City,
		State:  transaction.Address.Province,
		Zip:    transaction.Address.Zipcode,
	}

	creditCardModel := model.CreditCard{
		Type:   transaction.CreditCard.Visa,
		Name:   transaction.CreditCard.NameOfCard,
		Number: transaction.CreditCard.NumberCard,
		CVV:    transaction.CreditCard.CVV,
		Date:   transaction.CreditCard.ExpiredDate,
	}

	tx := r.DB.Model(&model.Address{}).Create(&addressModel)

	if tx.Error != nil {
		return tx.Error
	}

	tx = r.DB.Model(&model.CreditCard{}).Create(&creditCardModel)

	if tx.Error != nil {
		return tx.Error
	}

	transactionModel := model.Transaction{
		UserID:          transaction.UserID,
		AddressID:       addressModel.ID,
		CreditCardID:    creditCardModel.ID,
		GrandTotal:      transaction.GrandTotal,
		Status:          transaction.StatusTransaction,
		TransactionCode: transaction.TransactionCode,
	}

	tx = r.DB.Model(&model.Transaction{}).Create(&transactionModel)

	if tx.Error != nil {
		return tx.Error
	}

	var DetailTransactionModelList []model.TransactionDetail

	for _, transactionDetail := range transaction.TransactionDetail {
		DetailTransactionModelList = append(DetailTransactionModelList, model.TransactionDetail{
			TransactionID: transactionModel.ID,
			ProductID:     transactionDetail.ProductID,
			Qty:           transactionDetail.Qty,
			Subtotal:      transactionDetail.Subtotal,
		})
	}

	tx = r.DB.Model(&model.TransactionDetail{}).Create(&DetailTransactionModelList)
	if tx.Error != nil {
		return tx.Error
	}

	return
}

func (r *transactionRepo) Update(transaction entity.TransactionEntity) (err error) {
	tx := r.DB.Model(&model.Transaction{}).Where("id = ?", transaction.TransactionID).Where("user_id = ?", transaction.UserID).Update("status", transaction.StatusTransaction)
	if tx.Error != nil {
		return tx.Error
	}
	return
}

func (r *transactionRepo) DeleteCart(transaction entity.TransactionEntity) (err error) {
	tx := r.DB.Model(&model.Cart{}).Where("user_id = ?", transaction.UserID).Delete(&model.Cart{})
	if tx.Error != nil {
		return tx.Error
	}
	return
}

func (r *transactionRepo) FindAll(transaction entity.TransactionEntity) (result []entity.TransactionEntity, err error) {
	var transactionModelList []model.Transaction

	tx := r.DB.Model(&model.Transaction{}).Where("user_id", transaction.UserID).Where("status = ?", transaction.StatusTransaction).Find(&transactionModelList)
	if err != nil {
		return result, tx.Error
	}

	for _, transactionModel := range transactionModelList {
		result = append(result, entity.TransactionEntity{
			TransactionID:     transactionModel.ID,
			UserID:            transactionModel.UserID,
			TransactionCode:   transactionModel.TransactionCode,
			GrandTotal:        transactionModel.GrandTotal,
			StatusTransaction: transactionModel.Status,
			TransactionDate:   int64(transactionModel.CreatedAt.Unix()),
		})
	}
	return
}

func (r *transactionRepo) Find(transaction entity.TransactionEntity) (result entity.TransactionEntity, err error) {
	var transactionModel model.Transaction

	tx := r.DB.Model(&model.Transaction{}).Where("id = ?", transaction.TransactionID).Where("user_id", transaction.UserID).Preload("Address").Preload("CreditCard").Preload("TransactionDetail.Product").First(&transactionModel)

	if tx.Error != nil {
		return result, tx.Error
	}

	result = entity.TransactionEntity{
		TransactionID:     transactionModel.ID,
		UserID:            transactionModel.UserID,
		TransactionCode:   transactionModel.TransactionCode,
		StatusTransaction: transactionModel.Status,
		GrandTotal:        transactionModel.GrandTotal,
		TransactionDate:   int64(transactionModel.CreatedAt.Unix()),
		Address: entity.AddressEntity{
			AddressID: transactionModel.Address.ID,
			Street:    transactionModel.Address.Street,
			City:      transactionModel.Address.City,
			Province:  transactionModel.Address.State,
			Zipcode:   transactionModel.Address.Zip,
		},
		CreditCard: entity.CreditCardEntity{
			CreditCardID: transactionModel.CreditCard.ID,
			Visa:         transactionModel.CreditCard.Type,
			NameOfCard:   transactionModel.CreditCard.Name,
			NumberCard:   transactionModel.CreditCard.Number,
			CVV:          transactionModel.CreditCard.CVV,
			ExpiredDate:  transactionModel.CreditCard.Date,
		},
	}
	for _, transactionDetail := range transactionModel.TransactionDetail {
		result.TransactionDetail = append(result.TransactionDetail, entity.TransactionDetailEntity{
			TransactionDetailID: transaction.TransactionID,
			Qty:                 transactionDetail.Qty,
			Subtotal:            transactionDetail.Subtotal,
			ProductID:           transactionDetail.ProductID,
			ProductName:         transactionDetail.Product.Name,
			ProductPrice:        transactionDetail.Product.Price,
			ImageUrl:            transactionDetail.Product.ImageUrl,
		})
	}

	return result, nil
}

func (r *transactionRepo) FindCart(transaction entity.TransactionEntity) (result []entity.TransactionDetailEntity, err error) {
	var cartList []model.Cart

	tx := r.DB.Model(&model.Cart{}).Where("user_id", transaction.UserID).Find(&cartList)
	if tx.Error != nil {
		return result, nil
	}

	for _, cart := range cartList {
		result = append(result, entity.TransactionDetailEntity{
			Qty:       cart.Qty,
			ProductID: cart.ProductID,
			Subtotal:  cart.Subtotal,
		})
	}

	return result, nil
}

func (r *transactionRepo) FindLastInsertedId() (row int, err error) {
	var transactionModel model.Transaction
	tx := r.DB.Model(&model.Transaction{}).Order("id DESC").First(&transactionModel)
	if tx.Error != nil {
		return -1, err
	}

	if transactionModel.ID > 0 {
		return int(transactionModel.ID), nil
	}

	return 0, nil
}
