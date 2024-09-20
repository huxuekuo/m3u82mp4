package library

import "github.com/go-redis/redis/v8"

var RDB *redis.Client

func NewRedis() *redis.Client {
	if RDB != nil {
		return RDB
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     "152.136.33.217:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return RDB
}
