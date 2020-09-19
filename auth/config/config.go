package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// DB is global DB Config
var Config = map[string]string{
	"driver":    "postgres",
	"constring": "host=localhost port=5432 user=postgres dbname=tahor_auth password=123456 sslmode=disable",
}

var DB *gorm.DB

var Cachestore = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})
