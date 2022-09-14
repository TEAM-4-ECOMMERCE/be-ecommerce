package cartusecase

import (
	entity "e-commerce/domains/cart/entity"
)

type cartUseCase struct {
	Repo entity.IrepoCart
}

func New(repo entity.IrepoCart) *cartUseCase {
	return &cartUseCase{
		Repo: repo,
	}
}

func (u *cartUseCase) Store(cart entity.CartEntity) (err error) {
	cartResult, err := u.Repo.Find(cart)
	if cartResult.CartID > 0 {
		return
	}

	if err.Error() != "record not found" {
		return err
	}

	product, err := u.Repo.FindProduct(cart)
	if err != nil {
		return err
	}

	cart.Subtotal = product.Subtotal
	_, errInsert := u.Repo.Insert(cart)

	return errInsert
}

func (u *cartUseCase) Update(cart entity.CartEntity) (err error) {
	product, err := u.Repo.Find(cart)
	if err != nil {
		return err
	}
	cart.Subtotal = float64(cart.Qty) * product.ProductPrice

	_, err = u.Repo.Update(cart)
	return err
}

func (u *cartUseCase) Delete(cart entity.CartEntity) (err error) {
	_, err = u.Repo.Delete(cart)
	if err != nil {
		return err
	}

	return nil
}

func (u *cartUseCase) GetList(cart entity.CartEntity) (result []entity.CartEntity, err error) {
	carts, err := u.Repo.FindAll(cart)

	if len(carts) < 1 {
		return result, nil
	}

	if err != nil {
		return result, err
	}

	grandTotal := 0

	for _, cart := range carts {
		grandTotal += int(cart.Subtotal)
	}

	for _, cart := range carts {
		cart.GrandTotal = float64(grandTotal)
		cart.TotalOrderProduct = uint(len(carts))

		result = append(result, cart)
	}

	return
}
