package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gueldenstone/lighttracker/log"
)

func connectRedis(host, password string) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Infof("connected to redis on %s", host)
			return nil
		},
	})
	return client
}
