package delivery

import (
	reg "e-commerce/domains/register/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterControl struct{
	RegisterInterface reg.IregisterInterface
}

func (control *RegisterControl) CreateUser(c echo.Context) error {
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

func NewController(e *echo.Echo, logic reg.IregisterInterface) {
	controller := &RegisterControl{
		RegisterInterface: logic,
	}

	e.POST("/register", controller.CreateUser, middlewares.JWTMiddleware())

}