package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDoStopsAfterMaxAttempts(t *testing.T) {
	calls := 0
	want := errors.New("boom")
	err := Do(context.Background(), Policy{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond},
		func(context.Context) error { calls++; return want })
	if !errors.Is(err, want) || calls != 3 {
		t.Fatalf("err = %v, calls = %d", err, calls)
	}
}

func TestDoReturnsOnSuccess(t *testing.T) {
	calls := 0
	err := Do(context.Background(), Policy{MaxAttempts: 5, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond},
		func(context.Context) error {
			calls++
			if calls < 2 {
				return errors.New("not yet")
			}
			return nil
		})
	if err != nil || calls != 2 {
		t.Fatalf("err = %v, calls = %d", err, calls)
	}
}
