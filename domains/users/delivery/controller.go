package delivery

import (
	"e-commerce/domains/users/entity"
	"e-commerce/utils/helpers"
	"e-commerce/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type usercontrol struct {
	UserInterface users.IusecaseUser
}

func NewController(logic users.IusecaseUser) *usercontrol{
	return &usercontrol{
		UserInterface: logic,
	}
}

func (control *usercontrol) GetUser(c echo.Context) error {
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	results, err := control.UserInterface.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success", CoreList(results)))
}

func (control *usercontrol) UpdateUser (c echo.Context) error{
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("bad request to update data"))
	}
	_, err := control.UserInterface.UpdateUser(toCoreRequest(dataUpdate))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("fail to update data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("succses update data", userToken))
}

func (control *usercontrol) DeleteUser(c echo.Context) error{
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	var dataRemove UserRequest
	_, err := control.UserInterface.DeleteUser(toCoreRequest(dataRemove))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error delete user"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("delete succses", userToken))
}