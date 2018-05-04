package API

import (
	"fmt"
	"net/http"

	"app/controllers/Docker"
	"app/controllers/Redis"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type ConnectionData struct {
	URL string `json:"url" xml:"url"`
}

func Connection(c echo.Context) error {
	id := Docker.MakeContainer()
	key, err := Redis.SET(id)
	if err != nil {
		return c.String(500, "Redis is Panic")
	}

	res := &ConnectionData{
		URL: "http://localhost:1234/api/v1/judge/" + key,
	}
	return c.JSON(http.StatusOK, res)

}

func ConnectionTest(c echo.Context) error {
	res, err := Redis.GET(c.Param("key"))
	if err != nil {
		return c.String(404, "Not found ContainerID")
	}
	return c.String(200, res)
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
