package delivery

import (
	users "e-commerce/domains/users/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
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
	e.GET("/user", controller.GetUser, middlewares.JWTMiddleware())
	e.DELETE("/user", controller.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/user", controller.UpdateUser, middlewares.JWTMiddleware())

}

func (control *UserControl) GetUser(c echo.Context) error {
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}

	results, err := control.UserInterface.GetUser()

	userResult := UserResponse{}

	for _, user := range results {
		if user.UID == userToken {
			userResult.Id = userToken
			userResult.Email = user.Email
			userResult.Name = user.Name
			userResult.ImageUrl = user.ImageUrl
		}
	}

	if err != nil && userResult.Id < 1 {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success", userResult))
}

func (control *UserControl) UpdateUser(c echo.Context) error {
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	dataUpdate.UserID = userToken
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("bad request to update data"))
	}
	_, err := control.UserInterface.UpdateUser(toCoreRequest(dataUpdate))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("fail to update data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("succses update data"))
}

func (control *UserControl) DeleteUser(c echo.Context) error {
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("token nya tuan !"))
	}
	var dataRemove UserRequest
	dataRemove.UserID = userToken
	_, err := control.UserInterface.DeleteUser(toCoreRequest(dataRemove))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error delete user"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("delete succses"))
}
