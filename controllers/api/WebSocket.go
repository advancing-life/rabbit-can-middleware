package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	// "app/controllers/docker"
	// "app/controllers/redis"
	"github.com/advancing-life/rabbit-can-middleware/controllers/docker"
	"github.com/advancing-life/rabbit-can-middleware/controllers/redis"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type (
	// ConnectionData ...
	ConnectionData struct {
		URL         string `json:"url""`
		ContainerID string `json:"container_id"`
		RESULT      string `json:"result"`
	}
)

// GetMD5Hash ...
func GetMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Time.String(time.Now())))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Connection ...
func Connection(c echo.Context) error {
	key := GetMD5Hash()

	value, err := docker.Mk(key, c.Param("lang"))
	if err != nil {
		c.Logger().Error(err)
		return c.String(500, "Docker is Panic")
	}

	err = redis.SET(key, value)
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

// ExecutionEnvironment ...
func ExecutionEnvironment(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			execmd, err := receive(ws)
			if err != nil {
				ws.Close()
				c.Logger().Error(err)
			}

			exec := make(chan docker.ExecutionCommand)
			go docker.Exec(exec, execmd, c.Param("name"))

			for v := range exec {
				err := send(ws, v)
				if err != nil {
					ws.Close()
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func send(ws *websocket.Conn, send docker.ExecutionCommand) (err error) {
	err = websocket.JSON.Send(ws, send)
	fmt.Printf("Send data=\x1b[36m%#v\x1b[0m\n", send)
	return
}

func receive(ws *websocket.Conn) (rcv docker.ExecutionCommand, err error) {
	err = websocket.JSON.Receive(ws, &rcv)
	if err != nil {
		return
	}
	fmt.Printf("Receive data=\x1b[36m%#v\x1b[0m\n", rcv)
	return
}
