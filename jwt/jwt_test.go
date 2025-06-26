package jwt

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestNewJWT(t *testing.T) {
	redisCli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cli := New(Config{
		SecretKey:   "test",
		ExpireHours: 1,
		RedisPrefix: "test",
	}, redisCli)
	token, _ := cli.GenerateToken(123, "admin", "app")
	t.Log(token)
}
