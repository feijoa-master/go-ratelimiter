package limiter

import (
	"testing"
)

func TestTokenBucket(t *testing.T) {
	tb := NewTokenBucket(2, 1)

	if !tb.Allow() {
		t.Error("первый запрос должен быть разрешён")
	}

	if !tb.Allow() {
		t.Error("второй запрос должен быть разрешён")
	}

	if tb.Allow() {
		t.Error("третий запрос должен быть заблокирован")
	}

}
