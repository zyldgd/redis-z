package resource

import "github.com/go-redis/redis"

func GetRedisClient() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     "ubuntu:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return cli
}
