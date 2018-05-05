package Redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/k0kubun/pp"
)

func Init() (connection redis.Conn) {
	connection, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		pp.Println(err)
	}
	// defer connection.Close()
	return
}

func SET(key string, value string) (err error) {
	_, err = Init().Do("SET", key, value)
	if err != nil {
		pp.Println(err)
	}
	return
}

func GET(key string) (get string, err error) {
	get, err = redis.String(Init().Do("GET", key))
	if err != nil {
		pp.Println(err)
	}
	return
}

func EXISTS(key string) (exists bool) {
	exists, err := redis.Bool(Init().Do("EXISTS", key))
	if err != nil {
		pp.Println(err)
	}
	return
}
