package underlying

import (
	"github.com/gomodule/redigo/redis"
	"jun/utils/conf"
	"time"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     conf.GetGlobalConfig().Cache.MaxIdle,
		MaxActive:   conf.GetGlobalConfig().Cache.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				conf.GetGlobalConfig().Cache.Network,
				conf.GetGlobalConfig().Cache.RedisServer)
			if err != nil {
				return nil, err
			}
			auth := conf.GetGlobalConfig().Cache.Auth
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", conf.GetGlobalConfig().Cache.SelectNo); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func GetCache() redis.Conn {
	return rds.Get()
}

var (
	rds *redis.Pool
)

func init() {
	rds = newPool()
}
