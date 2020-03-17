package backoff

import (
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	b := New(10 * time.Second)
	for i := 0; i < 100; i++ {
		t.Log(b.Backoff(i))
	}
}
