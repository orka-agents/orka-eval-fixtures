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

func TestJitterWaitsWithinTheDelay(t *testing.T) {
	var waits []time.Duration
	old := randN
	randN = func(n time.Duration) time.Duration { waits = append(waits, n); return n / 2 }
	defer func() { randN = old }()
	_ = Do(context.Background(), Policy{MaxAttempts: 3, BaseDelay: 2 * time.Millisecond, MaxDelay: 3 * time.Millisecond, Jitter: true},
		func(context.Context) error { return errors.New("boom") })
	want := []time.Duration{2 * time.Millisecond, 3 * time.Millisecond}
	if len(waits) != len(want) || waits[0] != want[0] || waits[1] != want[1] {
		t.Fatalf("jitter bounds = %v, want %v", waits, want)
	}
}

func TestDoStopsWhenContextEnds(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	err := Do(ctx, Policy{MaxAttempts: 5, BaseDelay: time.Hour, MaxDelay: time.Hour, Jitter: true},
		func(context.Context) error { calls++; cancel(); return errors.New("boom") })
	if !errors.Is(err, context.Canceled) || calls != 1 {
		t.Fatalf("err = %v, calls = %d", err, calls)
	}
}
