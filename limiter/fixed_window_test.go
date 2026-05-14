package limiter

import (
	"testing"
	"time"
)

func TestFixedWindow_Reset(t *testing.T) {
	tb := NewFixedWindow(2, 100*time.Millisecond)

	tb.Allow()
	tb.Allow()

	// окно исчерпано
	if tb.Allow() {
		t.Error("третий запрос должен быть заблокирован")
	}

	// ждём сброса окна
	time.Sleep(150 * time.Millisecond)

	// теперь должен пройти
	if !tb.Allow() {
		t.Error("после сброса окна запрос должен быть разрешён")
	}
}
