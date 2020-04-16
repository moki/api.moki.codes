package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/moki/api.moki.codes/api"
	"github.com/moki/api.moki.codes/dotenv"
)

func main() {
	r := dotenv.NewReaderT(".env")
	r.Read()

	apiPort := os.Getenv("BACKEND_CONTAINER_PORT")
	redisPort := os.Getenv("REDIS_CONTAINER_PORT")
	redispw := os.Getenv("REDIS_PW")

	api := api.API{}
	api.Router = http.NewServeMux()
	api.RedisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"redis:"+redisPort,
				redis.DialPassword(redispw))
		},
	}
	api.Initialize()

	log.Fatal(
		http.ListenAndServeTLS(
			":"+apiPort,
			"certificates/_server.crt",
			"certificates/_server.key",
			api.Router,
		))
}
