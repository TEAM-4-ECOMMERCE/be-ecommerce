package routes

import (
	"e-commerce/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authhandler "e-commerce/domains/auth/handlers"
	authrepo "e-commerce/domains/auth/repositories"
	authusecase "e-commerce/domains/auth/usecases"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	/*
		Authentication Dependency
	*/
	authRepo := authrepo.New(db)
	authUsecase := authusecase.New(authRepo)
	authHandler := authhandler.New(authUsecase)

	/*
		Authentication Routes
	*/
	e.POST("/login", authHandler.Login)
}
