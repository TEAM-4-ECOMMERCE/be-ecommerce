package delivery

import (
	reg "e-commerce/domains/register/entity"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type registercontrol struct {
	RegisterInterface reg.IregisterInterface
}

func NewController(logic reg.IregisterInterface) *registercontrol {
	return &registercontrol{
		RegisterInterface: logic,
	}
}

func (control *registercontrol) CreateUser(c echo.Context) error {
	var dataRequest requestRegister

	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	if err := c.Validate(dataRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	_, err := control.RegisterInterface.CreateUser(toCoreRequest(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error create data"))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessGetResponseData("success create data"))
}
