package main

import (
	"fmt"
	"github.com/feijoa-master/go-ratelimiter/limiter"
	"github.com/feijoa-master/go-ratelimiter/middleware"
	"net/http"
)

func main() {
	tb := limiter.NewTokenBucket(5, 1)
	rl := middleware.New(tb)

	mux := http.NewServeMux()
	mux.Handle("/", rl.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})))

	http.ListenAndServe(":8081", mux)
}
