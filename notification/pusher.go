package notification

import (
	"github.com/go-redis/redis"
	"os"
	"fmt"
)

type Pusher struct {
	redisClient *redis.Client
}

func (self Pusher) Push(channel string, message string) error {
	err := self.redisClient.Publish(channel, message).Err()
	return err
}

func NewPusher() Pusher {
	pusher := &Pusher{}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pusher.redisClient = redisClient

	return *pusher
}