package limiter

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestFixedWindowRedis(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	client.Del(context.Background(), "test:fwindow")

	fw := NewFixedWindowRedis(client, "test:fwindow", 2, time.Second)

	if !fw.Allow() {
		t.Error("1й запрос должен пройти")
	}
	if !fw.Allow() {
		t.Error("2й запрос должен пройти")
	}
	if fw.Allow() {
		t.Error("3й запрос должен быть заблокирован")
	}

}
