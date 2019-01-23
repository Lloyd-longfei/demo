package database

import (
	"demo/conf"

	"github.com/garyburd/redigo/redis"
)

func RedisClient() (redis.Conn, error) {

	conf := conf.ReadConfFile()

	c, err := redis.Dial("tcp", conf.GetString("redis.host")+":"+conf.GetString("redis.port"))
	return c, err
}
