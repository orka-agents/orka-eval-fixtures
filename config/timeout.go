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
// number of seconds ("45"). Empty means DefaultTimeout. A value outside
// [MinTimeout, MaxTimeout] is clamped into that range, so a setting written
// for another service still yields a usable timeout.
func ParseTimeout(raw string) (time.Duration, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return DefaultTimeout, nil
	}
	d, err := parseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("config: invalid timeout %q", raw)
	}
	return min(max(d, MinTimeout), MaxTimeout), nil
}

func parseDuration(raw string) (time.Duration, error) {
	if seconds, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return time.Duration(seconds), nil
	}
	return time.ParseDuration(raw)
}
