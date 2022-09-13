package routes

import (
	"e-commerce/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authhandler "e-commerce/domains/auth/handlers"
	authrepo "e-commerce/domains/auth/repositories"
	authusecase "e-commerce/domains/auth/usecases"

	producthandler "e-commerce/domains/product/handlers"
	productrepo "e-commerce/domains/product/repositories"
	productusecase "e-commerce/domains/product/usecases"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	/*
		Dependency Injection
	*/
	authRepo := authrepo.New(db)
	authUsecase := authusecase.New(authRepo)
	authHandler := authhandler.New(authUsecase)

	productRepo := productrepo.New(db)
	productUsecase := productusecase.New(productRepo)
	productHandler := producthandler.New(productUsecase)

	/*
		Routes
	*/
	e.POST("/login", authHandler.Login)

	e.GET("/products", productHandler.ProductList)
	e.GET("/product/:id", productHandler.Product)
}
