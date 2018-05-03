package Redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/k0kubun/pp"
)

func Init() redis.Conn {
	connection, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		pp.Println(err)
	}
	// defer connection.Close()
	return connection
}

func SET(value string) error {
	_, err := Init().Do("SET", "Name", value)
	if err != nil {
		pp.Println(err)
		return err
	}
	return nil
}

// func GET(key string) interface{} {
func GET(key string) string {
	n, err := redis.String(Init().Do("GET", key))
	if err != nil {
		pp.Println(err)
	}
	return n
}

// func EXISTS() bool {
// }
