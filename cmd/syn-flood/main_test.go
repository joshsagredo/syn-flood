package main

import (
	"context"
	"testing"
	"time"
)

func TestMainProgram(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go func() {
		main()
	}()

	select {
	case <-time.After(120 * time.Second):
		t.Logf("overslept")
	case <-ctx.Done():
		t.Logf("ending flood")
	}
}
