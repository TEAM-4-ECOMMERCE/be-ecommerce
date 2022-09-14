package carthandler

import (
	entity "e-commerce/domains/cart/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	Usecase entity.IusecaseCart
}

func New(usecase entity.IusecaseCart) *cartHandler {
	return &cartHandler{
		Usecase: usecase,
	}
}

func (h *cartHandler) Store(c echo.Context) error {
	cartRequest := cartRequestStore{}

	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = c.Bind(&cartRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = h.Usecase.Store(entity.CartEntity{
		UserID:    uint(uid),
		ProductID: cartRequest.ProductID,
		Qty:       1,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success add to cart"))
}

func (h *cartHandler) Update(c echo.Context) error {
	cartId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	cartRequest := cartRequestUpdate{}
	err = c.Bind(&cartRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = h.Usecase.Update(entity.CartEntity{
		CartID: uint(cartId),
		Qty:    cartRequest.Qty,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success update to cart"))
}

func (h *cartHandler) Delete(c echo.Context) error {
	cartId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = h.Usecase.Delete(entity.CartEntity{
		CartID: uint(cartId),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success delete product cart"))
}

func (h *cartHandler) GetList(c echo.Context) error {
	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	carts, err := h.Usecase.GetList(entity.CartEntity{
		UserID: uint(uid),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	cartResponse := []cartResponse{}

	for _, cart := range carts {
		cartResponse = append(cartResponse, EntityToResponse(cart))
	}

	if len(cartResponse) < 1 {
		return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success get all data carts", map[string]interface{}{
			"total_order_product": 0,
			"grand_total":         0,
			"products":            cartResponse,
		}))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success get all data carts", map[string]interface{}{
		"total_order_product": carts[0].TotalOrderProduct,
		"grand_total":         carts[0].GrandTotal,
		"products":            cartResponse,
	}))
}
