package authhandler

import (
	entity "e-commerce/domains/auth/entity"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	Usecase entity.IusecaseAuth
}

func New(usecase entity.IusecaseAuth) *authHandler {
	return &authHandler{
		Usecase: usecase,
	}
}

func (h *authHandler) Login(c echo.Context) error {
	authRequest := authRequest{}

	if err := c.Validate(authRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	c.Bind(&authRequest)

	authEntity := requestToEntity(authRequest)
	token, err := h.Usecase.Login(authEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success logged", map[string]interface{}{
		"token": token,
	}))
}
