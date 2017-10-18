package route

import (
	. "github.com/JermineHu/ait/consts"
	 "github.com/JermineHu/ait/handler"
	"os"
	"strings"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type MeanErrorHandleType int

const (
	MeanErrorHandleTypeHeader MeanErrorHandleType = 1 + iota
	MeanErrorHandelTypeBody
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

type Mean struct {
	Echo *echo.Echo
}

func (m Mean) Engine() *Mean {

	// Global middleware
	m.Echo.Use(CheckDBSession)
	//m.Use(middleware.CSRF())
	m.Echo.Use(middleware.BodyLimit("2M"))
	m.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://b.daocloud.io", "https://b.daocloud.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	m.Echo.Use(middleware.Logger())
	m.Echo.Use(middleware.Recover())
	// new relic monitor
	//gorelic.InitNewrelicAgent("f243fdc54ca4b221bbabef85444e798a6d946335", "Berk", false)
	//m.Route.Use(gorelic.Handler)

	m.Echo.Use(InitHandler)
	m.Echo.Use(HeaderErrorHandler)
	m.Echo.Use(ErrorHandler)

	back := m.Echo.Group(BackendGroupRouteModuleName)
	{

		v1 := back.Group(GroupRouteVersion1Key)
		{
			//handler.LoginHandler(v1)
			if !strings.EqualFold(os.Getenv("IsDebug"), "true") {
				v1.Use(handler.AuthorizationHandler)
			}
			handler.WrapChatRoutes(v1)
			//handler.WrapUserRoutes(v1)
			//handler.WrapArticleRoutes(v1)
			//handler.WrapTagRoutes(v1)
			//handler.WrapFormRoutes(v1)
			//handler.WrapFeedbackRoutes(v1)
			//handler.WrapLotteryDrawRoutes(v1)
			//handler.WrapLotterySettingRoutes(v1)
		}

	}

	res := m.Echo.Group(ResourcesGroupRouteModuleName)
	{

		res.Static("/", os.Getenv(ResourcesPath))

	}

	return &m
}
