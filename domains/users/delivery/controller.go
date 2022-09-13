package delivery

import (
	"e-commerce/domains/users/entity"
	"e-commerce/utils/helpers"
	// "e-commerce/domains/users/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControl struct {
	UserInterface users.UserInterface
}

func NewController(e *echo.Echo, logic users.UserInterface) {
	controller := &UserControl{
		UserInterface: logic,
	}

	e.GET("/users",controller.GetUser)    //Need JWT
	e.GET("/users",controller.DeleteUser) //Need JWT


}


func (control *UserControl) GetUser(c echo.Context) error {
	results, err := control.UserInterface.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success", CoreList(results)))
}

func (control *UserControl) DeleteUser(c echo.Context) error{
	var dataRemove UserRequest
	_, err := control.UserInterface.DeleteUser(toCoreRequest(dataRemove))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error delete user"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("delete succses"))
}