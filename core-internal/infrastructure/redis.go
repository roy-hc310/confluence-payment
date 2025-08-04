package infrastructure

import (
	"confluence-payment/core-internal/utils"
	"time"

	"github.com/go-redis/redis"
	// "github.com/rs/zerolog/log"
)

type RedisInfra struct {
	Client *redis.Client
}

func NewRedisInfra() *RedisInfra {

	client := redis.NewClient(&redis.Options{
		Addr:        utils.GlobalEnv.RedisHost,
		Password:    utils.GlobalEnv.RedisPass,
		DB:          0,
		DialTimeout: time.Second * time.Duration(utils.GlobalEnv.RedisTimeOut),
		MaxRetries:  0,
	})

	return &RedisInfra{
		Client: client,
	}
}
