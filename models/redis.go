package models

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
)

var Pool *redis.Pool
var RPool *pools.ResourcePool

func InitRedis() {

	RPool = pools.NewResourcePool(func() (pools.Resource, error) {

		//local
		//c, err := redis.Dial("tcp", ":6379")

		//heroku
		c, err := redis.DialURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err.Error())
		}
		return ResourceConn{c}, err
	}, 20, 20, time.Minute)

	// Pool = &redis.Pool{
	// 	MaxIdle:   20,
	// 	MaxActive: 10000, // max number of connections
	// 	Dial: func() (redis.Conn, error) {
	//
	// 		//local
	// 		//c, err := redis.Dial("tcp", ":6379")
	//
	// 		//heroku
	// 		c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		return c, err
	// 	},
	// }
}
