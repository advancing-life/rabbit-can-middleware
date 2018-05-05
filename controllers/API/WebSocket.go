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

type CMD struct {
	ContainerID string `json:"container_id"`
	Command     string `json:"command"`
	Result      string `json:"result"`
}

func GetMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Time.String(time.Now())))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Connection(c echo.Context) error {
	key := GetMD5Hash()
	value, err := Docker.Mk(key, c.Param("lang"))
	if err != nil {
		return c.String(500, "Docker is Panic")
	}

	err = Redis.SET(key, value)
	if err != nil {
		return c.String(500, "Redis is Panic")
	}

	res := &ConnectionData{
		URL:    "ws://localhost:1234/api/v1/execution_environment/" + key,
		RESULT: value,
	}
	return c.JSON(http.StatusOK, res)

}

func ExecutionEnvironment(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		for {
			rcv := receive(ws)

			rcv.Result, _ = Docker.Exec(c.Param("name"), rcv.Command)
			send(ws, rcv)

			// _ = ws.Close()
		}
		// go receive(ws)
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func send(ws *websocket.Conn, send *CMD) {
	websocket.JSON.Send(ws, send)
	fmt.Printf("Send data=\x1b[36m%#v\x1b[0m\n", send)
}

func receive(ws *websocket.Conn) (rcv *CMD) {
	websocket.JSON.Receive(ws, &rcv)
	fmt.Printf("Receive data=\x1b[36m%#v\x1b[0m\n", rcv)
	return
}
