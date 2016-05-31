package models

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitRedis() {
	Pool = &redis.Pool{
		MaxIdle:   20,
		MaxActive: 20, // max number of connections
		Dial: func() (redis.Conn, error) {

			//local
			//c, err := redis.Dial("tcp", ":6379")

			//heroku
			c, err := redis.DialURL(os.Getenv("REDIS_URL"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
