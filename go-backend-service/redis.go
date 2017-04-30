package main

import "github.com/go-redis/redis"


var redisClient = redis.NewClient(&redis.Options{
	Addr:     "redis-master:" + getenv("REDIS_PORT", "6379"),
	Password: "",
	DB:       0,
})
