package routes


import (
	"github.com/labstack/echo/v4"
)

func HelloRoute(e *echo.Group) error {
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "Hello World")
	})
	return nil
}