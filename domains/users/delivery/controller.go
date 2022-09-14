package delivery

import (
	"e-commerce/domains/users/entity"
	"e-commerce/utils/helpers"
	"e-commerce/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControl struct {
	UserInterface users.IusecaseUser
}

func NewController(e *echo.Echo, logic users.IusecaseUser) {
	controller := &UserControl{
		UserInterface: logic,
	}
	e.GET("/users",controller.GetUser)    //Need JWT
	e.GET("/users",controller.DeleteUser, middlewares.JWTMiddleware()) //Need JWT
	e.POST("/users", controller.UpdateUser, middlewares.JWTMiddleware())

}


func (control *UserControl) GetUser(c echo.Context) error {
	results, err := control.UserInterface.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success", CoreList(results)))
}

func (control *UserControl) UpdateUser (c echo.Context) error{
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

func (control *UserControl) DeleteUser(c echo.Context) error{
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