package API

import (
	"fmt"
	"net/http"

	"app/controllers/Redis"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type ConnectionData struct {
	URL string `json:"url" xml:"url"`
}

func Connection(c echo.Context) error {
	err := Redis.SET("hoge") //test
	if err != nil {
		return c.String(500, "Error")
	}

	res := &ConnectionData{
		URL: "http://localhost:1234/api/v1/judge/" + Redis.GET("Name"),
	}
	return c.JSON(http.StatusOK, res)

}

func Judge(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
