# go-ratelimiter
A Go library for rate limiting HTTP services.
Supports three algorithms: Token Bucket, Fixed Window, and Sliding Window.
## Installation
go get github.com/feijoa-master/go-ratelimiter
## Usage
Пример кода
```
tb := limiter.NewTokenBucket(5, 1)
rl := middleware.New(tb)
mux := http.NewServeMux()
mux.Handle("/", rl.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "OK")
})))
http.ListenAndServe(":8081", mux)
```
## License
MIT
## Algorithms

- **Token Bucket** — allows bursts up to capacity, refills tokens at a fixed rate over time
- **Fixed Window** — counts requests in a fixed time window, resets counter when window expires
- **Sliding Window** — tracks exact timestamp of each request, provides smooth limiting without boundary spikes