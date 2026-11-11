// Package retry retries a function with capped exponential backoff.
package retry

import (
	"context"
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
}

// Do calls fn until it succeeds, the attempts run out or ctx ends. It returns
// the last error from fn, or ctx's error when ctx ends first. The waits are
// BaseDelay, 2*BaseDelay, 4*BaseDelay and so on, each capped at MaxDelay; there
// is no wait after the last attempt.
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
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
		delay = min(delay*2, p.MaxDelay)
	}
}
