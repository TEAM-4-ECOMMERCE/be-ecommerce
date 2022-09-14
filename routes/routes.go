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

	userscontrol "e-commerce/domains/users/delivery"
	usersrepo 	"e-commerce/domains/users/repositories"
	usersusecase "e-commerce/domains/users/usecase"

	registercontrol "e-commerce/domains/register/delivery"
	registerrepo "e-commerce/domains/register/repositories"
	registerusecase	"e-commerce/domains/register/usecase"

	categorycontrol "e-commerce/domains/category/delivery"
	categoryrepo	"e-commerce/domains/category/repositories"
	categoryusecase	"e-commerce/domains/category/usecase"
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

	usersRepo := usersrepo.NewDataBase(db)
	usersUsecase := usersusecase.NewLogic(usersRepo)
	usersControl := userscontrol.NewController(usersUsecase)

	registerRepo := registerrepo.NewRegister(db)
	registerUsecase := registerusecase.NewLogic(registerRepo)
	registerControl := registercontrol.NewController(registerUsecase)

	categoryRepo := categoryrepo.New(db)
	categoryUsecase := categoryusecase.New(categoryRepo)
	categoryControl := categorycontrol.New(categoryUsecase)

	/*
		Routes
	*/
	e.POST("/login", authHandler.Login)

	e.POST("/products", productHandler.Store, middlewares.JWTMiddleware())
	e.PUT("/product/:id", productHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/product/:id", productHandler.Delete, middlewares.JWTMiddleware())
	e.GET("/user/product", productHandler.ProductListUser, middlewares.JWTMiddleware())
	e.GET("/products", productHandler.ProductList)
	e.GET("/product/:id", productHandler.Product)

	e.GET("/carts", cartHandler.GetList, middlewares.JWTMiddleware())
	e.POST("/carts", cartHandler.Store, middlewares.JWTMiddleware())
	e.PUT("/cart/:id", cartHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", cartHandler.Delete, middlewares.JWTMiddleware())

	e.GET("/users", usersControl.GetUser,middlewares.JWTMiddleware())    
	e.GET("/users",usersControl.DeleteUser, middlewares.JWTMiddleware()) 
	e.POST("/users", usersControl.UpdateUser,middlewares.JWTMiddleware())

	e.POST("/register", registerControl.CreateUser,middlewares.JWTMiddleware())

	e.GET("/category", categoryControl.GetAllCategory,middlewares.JWTMiddleware())
}
