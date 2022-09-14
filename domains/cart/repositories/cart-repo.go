package cartrepo

import (
	entity "e-commerce/domains/cart/entity"
	model "e-commerce/domains/cart/models"
	"errors"

	"gorm.io/gorm"
)

type cartRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *cartRepo {
	return &cartRepo{
		DB: db,
	}
}

func (r *cartRepo) Insert(cart entity.CartEntity) (affectedRow int, err error) {
	cartValue := model.EntityToModel(cart)

	tx := r.DB.Model(&model.Cart{}).Create(&cartValue)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("fail to insert")
	}

	return int(tx.RowsAffected), nil
}

func (r *cartRepo) Update(cart entity.CartEntity) (affectedRow int, err error) {
	tx := r.DB.Model(&model.Cart{}).Where("id = ?", cart.CartID).Update("qty", cart.Qty)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("fail to update")
	}

	return int(tx.RowsAffected), nil
}

func (r *cartRepo) Delete(cart entity.CartEntity) (affectedRow int, err error) {
	cartValue := model.EntityToModel(cart)
	cartValue.ID = cart.CartID

	tx := r.DB.Model(&model.Cart{}).Delete(&cartValue)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("fail to update")
	}

	return int(tx.RowsAffected), nil
}

func (r *cartRepo) FindAll(cart entity.CartEntity) (result []entity.CartEntity, err error) {
	cartModel := []model.Cart{}

	tx := r.DB.Model(&model.Cart{}).Preload("Product").Where("user_id", cart.UserID).Find(&cartModel)

	if tx.Error != nil {
		return result, tx.Error
	}

	for _, cart := range cartModel {
		result = append(result, model.ModelToEntity(cart))
	}

	return result, nil
}

func (r *cartRepo) Find(cart entity.CartEntity) (result entity.CartEntity, err error) {
	cartModel := model.Cart{}
	cartModel.ID = cart.CartID

	tx := r.DB.Model(&model.Cart{}).Preload("Product").Where("user_id", cart.UserID).First(&cartModel)

	if tx.Error != nil {
		return result, tx.Error
	}

	return model.ModelToEntity(cartModel), nil
}
