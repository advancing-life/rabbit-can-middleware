package Redis

import (
	"crypto/md5"
	"encoding/hex"
	"time"

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

func GetMD5Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Time.String(time.Now())))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SET(value string) (Key string, err error) {
	Key = GetMD5Hash()
	_, err = Init().Do("SET", Key, value)
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
