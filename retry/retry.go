// Package retry retries a function with capped exponential backoff.
package retry

import (
	"context"
	"math/rand/v2"
	"time"
)

// Policy bounds a retry loop.
type Policy struct {
	// MaxAttempts is the total number of calls, including the first.
	MaxAttempts int
	// BaseDelay is the wait after the first failure. It doubles after every
	// further failure.
	BaseDelay time.Duration
	// MaxDelay caps the wait between attempts.
	MaxDelay time.Duration
	// Jitter, when true, waits a random duration in [0, delay) instead of the
	// full delay, so many clients retrying at once spread out.
	Jitter bool
}

// randN returns a random duration in [0, n). Tests replace it.
var randN = func(n time.Duration) time.Duration { return rand.N(n) }

// Do calls fn until it succeeds, the attempts run out or ctx ends. It returns
// the last error from fn, or ctx's error when ctx ends first.
func Do(ctx context.Context, p Policy, fn func(context.Context) error) error {
	attempts := max(p.MaxAttempts, 1)
	delay := p.BaseDelay
	var err error
	for attempt := 1; ; attempt++ {
		if err = fn(ctx); err == nil {
			return nil
		}
		if attempt == attempts {
			return err
		}
		wait := delay
		if p.Jitter && wait > 0 {
			wait = randN(wait)
		}
		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
		delay = min(delay*2, p.MaxDelay)
	}
}
