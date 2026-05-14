package limiter

import (
	"sync"
	"time"
)

type FixedWindow struct {
	limit       int
	count       int
	windowSize  time.Duration
	windowStart time.Time
	mu          sync.Mutex
}

func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:       limit,
		count:       0,
		windowSize:  windowSize,
		windowStart: time.Now(),
	}
}

func (f *FixedWindow) Allow() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now()
	if now.Sub(f.windowStart) >= f.windowSize {
		f.count = 0
		f.windowStart = now
	}
	if f.count < f.limit {
		f.count++
		return true
	}
	return false
}
