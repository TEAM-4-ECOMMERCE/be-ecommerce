package transactionhandlers

import (
	entity "e-commerce/domains/transaction/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	Usecase entity.IusecaseTransaction
}

func New(usecase entity.IusecaseTransaction) *transactionHandler {
	return &transactionHandler{
		Usecase: usecase,
	}
}

func (h *transactionHandler) Store(c echo.Context) error {
	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.FailedResponse(err.Error()))
	}

	var transactionRequest StoreRequest

	if err := c.Validate(transactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = c.Bind(&transactionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	transactionEntity := RequestToEntity(transactionRequest)
	transactionEntity.UserID = uint(uid)

	err = h.Usecase.Store(transactionEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success create transaction"))
}

func (h *transactionHandler) Update(c echo.Context) error {
	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.FailedResponse(err.Error()))
	}

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	var transactionRequest UpdateRequest

	if err := c.Validate(transactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	err = c.Bind(&transactionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	transactionEntity := entity.TransactionEntity{
		UserID:            uint(uid),
		TransactionID:     uint(transactionID),
		StatusTransaction: transactionRequest.Status,
	}

	err = h.Usecase.Update(transactionEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success update transaction"))
}

func (h *transactionHandler) GetSingle(c echo.Context) error {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.FailedResponse(err.Error()))
	}

	transactionEntity := entity.TransactionEntity{
		UserID:        uint(uid),
		TransactionID: uint(transactionID),
	}
	transaction, err := h.Usecase.GetSingle(transactionEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success get transaction", EntityToResponseSingle(transaction)))
}

func (h *transactionHandler) GetList(c echo.Context) error {
	status := c.QueryParam("status")

	uid, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.FailedResponse(err.Error()))
	}

	transactionEntity := entity.TransactionEntity{
		UserID:            uint(uid),
		StatusTransaction: status,
	}

	transactionList, err := h.Usecase.GetList(transactionEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success get all transaction", EntityToResponseList(transactionList)))
}
