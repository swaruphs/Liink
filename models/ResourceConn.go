package models

import "github.com/garyburd/redigo/redis"

type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}
