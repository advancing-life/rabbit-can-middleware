package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Init ...
func Init() (connection redis.Conn, err error) {
	connection, err = redis.Dial("tcp", "redis:6379")
	// defer connection.Close()
	return
}

// SET ...
func SET(key string, value string) (err error) {
	r, err := Init()
	if err != nil {
		return
	}

	_, err = r.Do("SET", key, value)
	if err != nil {
		return
	}
	return
}

// GET ...
func GET(key string) (get string, err error) {
	r, err := Init()
	if err != nil {
		return
	}
	get, err = redis.String(r.Do("GET", key))
	if err != nil {
		return
	}
	return
}

// EXISTS ...
func EXISTS(key string) (exists bool, err error) {
	r, err := Init()
	if err != nil {
		return
	}
	exists, err = redis.Bool(r.Do("EXISTS", key))
	if err != nil {
		return
	}
	return
}
