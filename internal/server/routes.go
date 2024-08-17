package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", HelloWorldHandler)

	// Register user routes
	userGroup := e.Group("/users")
	userGroup.GET("/", s.userHandler.ListUser)
	userGroup.POST("/", s.userHandler.CreateUser)
	userGroup.PATCH("/:id", s.userHandler.UpdateUser)

	return e
}

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
