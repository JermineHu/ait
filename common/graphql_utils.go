package common

import (
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
)

func SetGraphqlHandler(e *echo.Echo,path string,h *handler.Handler)  {

	e.GET(path, func(c echo.Context) error {

		h.ContextHandler(c.Request().Context(),c.Response().Writer,c.Request())

		return nil
	})
	e.POST(path, func(c echo.Context) error {

		h.ContextHandler(c.Request().Context(),c.Response().Writer,c.Request())

		return nil
	})
}
