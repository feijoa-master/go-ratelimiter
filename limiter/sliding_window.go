package limiter

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	requests   []time.Time
	limit      int
	windowSize time.Duration
	mu         sync.Mutex
}

func NewSlidingWindow(limit int, windowSize time.Duration) *SlidingWindow {
	return &SlidingWindow{
		requests:   make([]time.Time, 0),
		limit:      limit,
		windowSize: windowSize,
	}
}

func (s *SlidingWindow) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	filtered := s.requests[:0]
	for _, t := range s.requests {
		if now.Sub(t) <= s.windowSize {
			filtered = append(filtered, t)
		}
	}
	s.requests = filtered

	if len(s.requests) < s.limit {
		s.requests = append(s.requests, now)
		return true
	}
	return false

}
