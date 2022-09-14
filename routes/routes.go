package routes

import (
	"e-commerce/config"
	"e-commerce/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authhandler "e-commerce/domains/auth/handlers"
	authrepo "e-commerce/domains/auth/repositories"
	authusecase "e-commerce/domains/auth/usecases"

	producthandler "e-commerce/domains/product/handlers"
	productrepo "e-commerce/domains/product/repositories"
	productusecase "e-commerce/domains/product/usecases"

	carthandler "e-commerce/domains/cart/handlers"
	cartrepo "e-commerce/domains/cart/repositories"
	cartusecase "e-commerce/domains/cart/usecases"

	registerrepo "e-commerce/domains/register/data"
	registerhandler "e-commerce/domains/register/delivery"
	registerusecase "e-commerce/domains/register/usecase"

	usersrepo "e-commerce/domains/users/data"
	usershandler "e-commerce/domains/users/delivery"
	usersusecase "e-commerce/domains/users/usecase"
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

	cartRepo := cartrepo.New(db)
	cartUsecase := cartusecase.New(cartRepo)
	cartHandler := carthandler.New(cartUsecase)

	registerRepo := registerrepo.NewRegister(db)
	registerUsecase := registerusecase.NewLogic(registerRepo)
	registerhandler.NewController(e, registerUsecase)

	usersRepo := usersrepo.NewDataBase(db)
	usersUsecase := usersusecase.NewLogic(usersRepo)
	usershandler.NewController(e, usersUsecase)

	/*
		Routes
	*/
	e.POST("/login", authHandler.Login)

	e.GET("/products", productHandler.ProductList)
	e.GET("/product/:id", productHandler.Product)

	e.GET("/carts", cartHandler.GetList, middlewares.JWTMiddleware())
	e.POST("/carts", cartHandler.Store, middlewares.JWTMiddleware())
	e.PUT("/cart/:id", cartHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", cartHandler.Delete, middlewares.JWTMiddleware())
}
