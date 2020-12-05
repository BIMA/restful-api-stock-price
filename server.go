package main

import (
	"github.com/kaabima/restful-api-stock-price/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!\n")
	})
	e.POST("/price", controllers.GrabPrice)

	// Run server
	e.Logger.Fatal(e.Start(":9000"))
}
