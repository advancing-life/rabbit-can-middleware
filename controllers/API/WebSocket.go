package API

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"app/controllers/Docker"
	"app/controllers/Redis"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type ConnectionData struct {
	URL    string `json:"url" xml:"url"`
	RESULT string `json:"result" xml:"result"`
}

func GetMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Time.String(time.Now())))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Connection(c echo.Context) error {
	key := GetMD5Hash()
	value, err := Docker.MakeContainer(key, c.Param("lang"))
	if err != nil {
		return c.String(500, "Docker is Panic")
	}

	err = Redis.SET(key, value)
	if err != nil {
		return c.String(500, "Redis is Panic")
	}

	res := &ConnectionData{
		URL:    "http://localhost:1234/api/v1/judge/" + key,
		RESULT: value,
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
