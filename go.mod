module github.com/shareed2k/echo_limiter

go 1.14

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/labstack/echo/v4 v4.9.0
	github.com/shareed2k/go_limiter v0.0.7
)

replace github.com/shareed2k/go_limiter => github.com/palicao/go_limiter v0.0.8
