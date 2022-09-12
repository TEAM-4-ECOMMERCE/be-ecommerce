package main

import (
	"e-commerce/config"

	database "e-commerce/utils/database/gorm-mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.GetConfig()
	database.InitDB(cfg)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	err := e.Start(":" + cfg.SERVER_PORT)

	if err != nil {
		panic(err)
	}
}
