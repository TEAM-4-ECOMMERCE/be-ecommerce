package productrepo

import (
	entity "e-commerce/domains/product/entity"
	productModel "e-commerce/domains/product/models"
	"errors"

	"gorm.io/gorm"
)

type productRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *productRepo {
	return &productRepo{
		DB: db,
	}
}

func (r *productRepo) Insert(product entity.ProductEntity) (affectedRow int, err error) {
	entityModel := productModel.ProductEntityToModel(product)

	tx := r.DB.Model(&productModel.Product{}).Create(&entityModel)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("data not inserted")
	}

	return int(tx.RowsAffected), nil
}

func (r *productRepo) Update(product entity.ProductEntity) (affectedRow int, err error) {
	entityModel := productModel.ProductEntityToModel(product)

	tx := r.DB.Model(&productModel.Product{}).Where("id = ?", product.ProductID).Where("user_id = ?", product.UserID).Updates(&entityModel)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("data not updated")
	}

	return int(tx.RowsAffected), nil
}

func (r *productRepo) Delete(product entity.ProductEntity) (affectedRow int, err error) {
	entityModel := productModel.ProductEntityToModel(product)
	entityModel.ID = product.ProductID
	tx := r.DB.Model(&productModel.Product{}).Where("user_id = ?", product.UserID).Delete(&entityModel)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected < 1 {
		return int(tx.RowsAffected), errors.New("data not deleted")
	}

	return int(tx.RowsAffected), nil
}

func (r *productRepo) FindAll(product entity.ProductEntity) (result []entity.ProductEntity, err error) {
	productModels := []productModel.Product{}

	tx := r.DB.Model(&productModel.Product{})

	if product.UserID > 0 {
		tx.Where("user_id = ?", product.UserID)
	}

	if product.SearchQuery != "" {
		tx.Where("name LIKE ?", "%"+product.SearchQuery+"%")
	}

	if product.CategoryID > 0 {
		tx.Where("category_id = ?", product.CategoryID)
	}

	tx.Preload("Category").Limit(int(product.PageSize)).Offset(int(product.CurrentPage)).Find(&productModels)

	if tx.Error != nil {
		return result, tx.Error
	}

	for _, product := range productModels {
		result = append(result, productModel.ProductModelToEntity(product))
	}

	return result, nil
}

func (r *productRepo) Find(product entity.ProductEntity) (result entity.ProductEntity, err error) {
	modelProduct := productModel.Product{}
	modelProduct.ID = product.ProductID

	tx := r.DB.Preload("Category").First(&modelProduct)

	if tx.Error != nil {
		return result, tx.Error
	}

	return productModel.ProductModelToEntity(modelProduct), nil
}

func (r *productRepo) CountData(uid uint) (totalRow uint, err error) {
	modelProduct := []productModel.Product{}

	tx := r.DB
	if uid > 0 {
		tx.Where("user_id = ?", uid)
	}

	tx.Find(&modelProduct)

	if tx.Error != nil {
		return totalRow, tx.Error
	}

	totalRow = uint(len(modelProduct))

	return totalRow, nil
}
