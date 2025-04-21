package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/charmbracelet/log"
)

func HelloMW(next echo.HandlerFunc) echo.HandlerFunc {
	log.Info("Hello Middleware")
	return func(c echo.Context) error {
		return next(c)
	}
}