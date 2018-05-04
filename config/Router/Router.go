package Router

import (
	"app/controllers/API"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Route
	v1 := e.Group("api/v1")
	{
		v1.GET("/", API.Index)
		v1.GET("/connection", API.Connection)
		v1.GET("/connection_test/:key", API.ConnectionTest)
	}

	return e
}
