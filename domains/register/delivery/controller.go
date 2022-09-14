package delivery

import (
	reg "e-commerce/domains/register/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type registercontrol struct{
	RegisterInterface reg.IregisterInterface
}

func NewController(logic reg.IregisterInterface) *registercontrol{
	return &registercontrol{
		RegisterInterface: logic,
	}
}

func (control *registercontrol) CreateUser(c echo.Context) error {
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	var dataRequest requestRegister
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}
	_, err := control.RegisterInterface.CreateUser(toCoreRequest(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error create data"))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessGetResponse("success create data", userToken))
}
