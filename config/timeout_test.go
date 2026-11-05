package config

import (
	"errors"
	"testing"
	"time"
)

func TestParseTimeout(t *testing.T) {
	tests := []struct {
		in      string
		want    time.Duration
		wantErr bool
	}{
		{in: "", want: DefaultTimeout},
		{in: "45", want: 45 * time.Second},
		{in: "45s", want: 45 * time.Second},
		{in: "2m", want: MaxTimeout},
		{in: "500ms", wantErr: true},
		{in: "soon", wantErr: true},
	}
	for _, tt := range tests {
		got, err := ParseTimeout(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("ParseTimeout(%q) error = %v", tt.in, err)
		}
		if err != nil && !errors.Is(err, ErrInvalidTimeout) {
			t.Fatalf("ParseTimeout(%q) error %v does not wrap ErrInvalidTimeout", tt.in, err)
		}
		if got != tt.want {
			t.Fatalf("ParseTimeout(%q) = %s, want %s", tt.in, got, tt.want)
		}
	}
}
