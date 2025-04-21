package routes

import (
	"fmt"
	"strings"
	"github.com/joho/godotenv"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/charmbracelet/log"
)

func Router(port int16) {
	godotenv.Load(".env")
	usernameEnv := os.Getenv("DOCK_USERNAME")
	log.Info(usernameEnv)
  	passwordEnv := os.Getenv("DOCK_PASSWORD")
	log.Info(passwordEnv)
  e := echo.New()
  e.Use(middleware.BasicAuth(
	func(username, password string, c echo.Context) (bool, error) {
		if username == usernameEnv && password == passwordEnv {
			return true, nil
		}
		return false, nil
	},
  ))
  // serve everything under frontend/dist except /api*
  e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
    // skip static for /api and /api/*
    Skipper: func(c echo.Context) bool {
      p := c.Request().URL.Path
      return p == "/api" || strings.HasPrefix(p, "/api/")
    },
    Root:  "frontend/build",
    Index: "index.html",
    HTML5: true,
  }))

  api := e.Group("/api")
  HelloRoute(api)
  SystemResources(api)
  e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
