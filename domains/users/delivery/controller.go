package delivery

import (
	users "e-commerce/domains/users/entity"
	"e-commerce/middlewares"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type usercontrol struct {
	UserInterface users.IusecaseUser
}

func NewController(logic users.IusecaseUser) *usercontrol {
	return &usercontrol{
		UserInterface: logic,
	}
}

func (control *usercontrol) GetUser(c echo.Context) error {
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

	if err != nil || userResult.Id < 1 {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success", userResult))
}

func (control *usercontrol) UpdateUser(c echo.Context) error {
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

	if err := c.Validate(dataUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	_, err := control.UserInterface.UpdateUser(toCoreRequest(dataUpdate))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("fail to update data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("succses update data"))
}

func (control *usercontrol) DeleteUser(c echo.Context) error {
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
