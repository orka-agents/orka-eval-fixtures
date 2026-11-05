// Package config parses settings.
package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Timeout bounds.
const (
	DefaultTimeout = 30 * time.Second
	MinTimeout     = time.Second
	MaxTimeout     = 60 * time.Second
)

// ParseTimeout parses a timeout written as a Go duration ("45s") or a whole
// number of seconds ("45"). Empty means DefaultTimeout. A value above
// MaxTimeout is clamped to MaxTimeout; a value below MinTimeout is an error.
func ParseTimeout(raw string) (time.Duration, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return DefaultTimeout, nil
	}
	var d time.Duration
	if seconds, err := strconv.Atoi(raw); err == nil {
		d = time.Duration(seconds) * time.Second
	} else {
		d, err = time.ParseDuration(raw)
		if err != nil {
			return 0, fmt.Errorf("config: invalid timeout %q", raw)
		}
	}
	if d < MinTimeout {
		return 0, fmt.Errorf("config: timeout %s is below %s", d, MinTimeout)
	}
	return min(d, MaxTimeout), nil
}
