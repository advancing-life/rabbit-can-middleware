package API

import (
    "github.com/labstack/echo"
    "net/http"
)


func Index(c echo.Context) error {
    return c.String(http.StatusOK, "IndexRequest")
}
