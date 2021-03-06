package router

import (
	"github.com/advancing-life/rabbit-can-middleware/controllers/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init ...
func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, //AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Route
	v1 := e.Group("api/v1")
	{
		v1.GET("/", api.Index)
		v1.GET("/connection/:lang", api.Connection)
		v1.GET("/execution_environment/:name", api.ExecutionEnvironment)
	}

	return e
}
