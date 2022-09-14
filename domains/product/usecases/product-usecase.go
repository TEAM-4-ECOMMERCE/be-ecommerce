package productusecase

import (
	entity "e-commerce/domains/product/entity"
)

type productUsecase struct {
	Repo entity.IrepoProduct
}

func New(repo entity.IrepoProduct) *productUsecase {
	return &productUsecase{
		Repo: repo,
	}
}

func (u *productUsecase) Store(product entity.ProductEntity) (err error) {
	_, err = u.Repo.Insert(product)

	if err != nil {
		return err
	}

	return nil
}

func (u *productUsecase) Update(product entity.ProductEntity) (err error) {
	_, err = u.Repo.Update(product)

	if err != nil {
		return err
	}

	return nil
}

func (u *productUsecase) Delete(product entity.ProductEntity) (err error) {
	_, err = u.Repo.Delete(product)

	if err != nil {
		return err
	}

	return nil
}

func (u *productUsecase) GetList(product entity.ProductEntity) (result []entity.ProductEntity, err error) {
	if product.CurrentPage <= 0 {
		product.CurrentPage = 1
	}

	offset := (product.CurrentPage - 1) * product.PageSize
	product.CurrentPage = offset

	result, err = u.Repo.FindAll(product)

	return result, err
}

func (u *productUsecase) GetSingle(product entity.ProductEntity) (result entity.ProductEntity, err error) {
	result, err = u.Repo.Find(product)
	return result, err
}
