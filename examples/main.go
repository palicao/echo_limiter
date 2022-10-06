package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	limiter "github.com/shareed2k/echo_limiter"
)

func main() {
	e := echo.New()

	option, err := redis.ParseURL("redis://127.0.0.1:6379/0")
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(option)
	_ = client.FlushDB(context.Background()).Err()

	// 3 requests per 10 seconds max
	cfg := limiter.Config{
		Rediser:   client,
		Max:       3,
		Burst:     3,
		Period:    10 * time.Second,
		Algorithm: limiter.GCRAAlgorithm,
	}

	e.Use(limiter.NewWithConfig(cfg))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
