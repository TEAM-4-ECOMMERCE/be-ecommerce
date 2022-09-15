package delivery

import (
	categoryentity "e-commerce/domains/category/entity"
	"e-commerce/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryControl struct {
	categoryUsecase categoryentity.IusecaseCategory
}

func New(logic categoryentity.IusecaseCategory) *CategoryControl {
	return &CategoryControl{
		categoryUsecase: logic,
	}
}

func (control *CategoryControl) GetAllCategory(c echo.Context) error {
	result, err := control.categoryUsecase.GetCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("internal server error"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("succses get category", CoreList(result)))
}
