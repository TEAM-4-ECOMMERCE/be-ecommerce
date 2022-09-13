package producthandler

import (
	"e-commerce/config"
	entity "e-commerce/domains/product/entity"
	"e-commerce/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type producthandler struct {
	Usecase entity.IusecaseProduct
}

func New(usecase entity.IusecaseProduct) *producthandler {
	return &producthandler{
		Usecase: usecase,
	}
}

func (h *producthandler) ProductList(c echo.Context) error {
	cfg := config.GetConfig()
	var currentPage int
	var pageSize int
	var categoryId int
	var err error

	if c.QueryParam("current_page") != "" {
		currentPage, err = strconv.Atoi(c.QueryParam("current_page"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
	}

	if c.QueryParam("page_size") != "" {
		pageSize, err = strconv.Atoi(c.QueryParam("page_size"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
	}

	if c.QueryParam("category_id") != "" {
		categoryId, err = strconv.Atoi(c.QueryParam("category_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
	}

	searchQuery := c.QueryParam("q")

	productEntity := entity.ProductEntity{
		CurrentPage: uint(currentPage),
		PageSize:    uint(pageSize),
		CategoryID:  uint(categoryId),
		SearchQuery: searchQuery,
	}

	products, err := h.Usecase.GetList(productEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	var productResponse []ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, EntityToProductResponse(product))
	}

	stringPreviousPage := strconv.Itoa(currentPage - 1)
	stringNextPage := strconv.Itoa(currentPage + 1)
	stringPageSize := strconv.Itoa(pageSize)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":       true,
		"message":       "success get product list",
		"previous_page": cfg.BASE_URL + "/products?current_page=" + stringPreviousPage + "&page_size=" + stringPageSize,
		"current_page":  currentPage,
		"next_page":     cfg.BASE_URL + "/products?current_page=" + stringNextPage + "&page_size=" + stringPageSize,
		"data":          productResponse,
	})
}

func (h *producthandler) Product(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	productEntity := entity.ProductEntity{
		ProductID: uint(productId),
	}

	product, err := h.Usecase.GetSingle(productEntity)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success get product", EntityToProductResponse(product)))
}
