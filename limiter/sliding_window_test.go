package limiter

import (
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	tb := NewSlidingWindow(2, 100*time.Millisecond)

	if !tb.Allow() {
		t.Error("1й запрос должен пройти")
	}

	if !tb.Allow() {
		t.Error("2й запрос должен пройти")
	}

	if tb.Allow() {
		t.Error("3rd запрос должен быть заблокирован")
	}

	time.Sleep(110 * time.Millisecond)

	if !tb.Allow() {
		t.Error("4th запрос должен пройти")
	}

}
