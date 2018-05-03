package Router

import (
	"net/http"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
    e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.GET("/", index)
    return e
}

func index(c echo.Context) error {
    return c.String(http.StatusOK, "indexRequest")
}
