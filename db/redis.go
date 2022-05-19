/**
  @author: qianyi  2022/5/17 20:03:00
  @note:
*/
package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"goHomework/config"
)

var (
	Client *redis.Client
)

func InitRedis(cfg *config.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err = Client.Ping().Result()
	if err != nil {
		return
	}
	return nil
}
