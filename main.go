package main

import (
	"e-commerce/config"

	database "e-commerce/utils/database/gorm-mysql"

	"e-commerce/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	routes.InitRoutes(e, db, cfg)

	err := e.Start(":" + cfg.SERVER_PORT)

	if err != nil {
		panic(err)
	}
}
