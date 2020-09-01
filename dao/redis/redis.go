package redis

import (
	"fmt"

	"tanjunchen.io.webapp/setting"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
)

// 初始化连接
func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		// Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Addr: fmt.Sprintf("%s:%d", setting.Conf.RedisConfig.Host, setting.Conf.RedisConfig.Port),
		DB:   viper.GetInt("redis.db"), // use default DB
	})

	_, err = Client.Ping().Result()
	return
}
