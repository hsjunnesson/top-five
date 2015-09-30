package main

import (
	"os"
	"net/url"
	"log"
	
	"github.com/hsjunnesson/top-five/Godeps/_workspace/src/gopkg.in/redis.v3"
)

var redisClient = OpenRedisClient()

func OpenRedisClient() *redis.Client {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Panic("$REDIS_URL not set")
	}
	
	server, password := ParseRedisUrl(redisUrl)

	client := redis.NewClient(&redis.Options{
		Addr: server,
		Password: password,
		DB: 0,
	})
	
	return client
}

func ParseRedisUrl(redisUrl string) (string, string) {
	redisInfo, err := url.Parse(redisUrl)
	if err != nil {
		log.Panic("Couldn't parse redis url", err)
	}
	
	server := redisInfo.Host
	password := ""
	if redisInfo.User != nil {
		password, _ = redisInfo.User.Password()
	}
	return server, password
}

