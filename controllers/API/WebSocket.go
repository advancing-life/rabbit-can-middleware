package API

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"app/controllers/Docker"
	"app/controllers/Redis"
	// "github.com/advancing-life/rabbit-can-middleware/controllers/Docker"
	// "github.com/advancing-life/rabbit-can-middleware/controllers/Redis"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type (
	ConnectionData struct {
		URL         string `json:"url""`
		ContainerID string `json:"container_id"`
		RESULT      string `json:"result"`
	}
)

func GetMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Time.String(time.Now())))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Connection(c echo.Context) error {
	key := GetMD5Hash()

	value, err := Docker.Mk(key, c.Param("lang"))
	if err != nil {
		c.Logger().Error(err)
		return c.String(500, "Docker is Panic")
	}

	err = Redis.SET(key, value)
	if err != nil {
		c.Logger().Error(err)
		return c.String(500, "Redis is Panic")
	}

	res := &ConnectionData{
		URL:         "ws://localhost:1234/api/v1/execution_environment/" + key,
		ContainerID: key,
		RESULT:      value,
	}
	return c.JSON(http.StatusOK, res)

}

func ExecutionEnvironment(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			excmd, err := receive(ws)
			if err != nil {
				c.Logger().Error(err)
			}

			ch := make(chan Docker.ExecutionCommand)
			go Docker.Exec(ch, c.Param("name"), excmd.Command)

			for v := range ch {
				send(ws, v)
			}

			// excmd.Result, excmd.ExitStatus, _ = Docker.Exec(c.Param("name"), excmd.Command)

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func send(ws *websocket.Conn, send Docker.ExecutionCommand) {
	websocket.JSON.Send(ws, send)
	fmt.Printf("Send data=\x1b[36m%#v\x1b[0m\n", send)
}

func receive(ws *websocket.Conn) (rcv Docker.ExecutionCommand, err error) {
	err = websocket.JSON.Receive(ws, &rcv)
	if err != nil {
		return
	}
	fmt.Printf("Receive data=\x1b[36m%#v\x1b[0m\n", rcv)
	return
}
