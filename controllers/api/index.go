package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// Responce ...
type Responce struct {
	Status  int    `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}

// Index ...
func Index(c echo.Context) error {
	res := &Responce{
		Status:  200,
		Message: "ヽ（　＾ω＾）ﾉｻｸｾｽ！",
	}
	return c.JSON(http.StatusOK, res)
}
