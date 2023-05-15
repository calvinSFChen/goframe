package redis

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var redisConnect *gredis.Redis

func init() {
	redisConnect = g.Redis()
}

func Connect() *gredis.Redis {
	if redisConnect == nil {
		return g.Redis()
	}
	return redisConnect
}
