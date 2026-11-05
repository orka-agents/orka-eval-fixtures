package textutil

import (
	"strings"
	"testing"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		in   string
		max  int
		want string
	}{
		{in: "hello", max: 10, want: "hello"},
		{in: "hello", max: 5, want: "hello"},
		{in: "hello world", max: 8, want: "hello…"},
		{in: "abc", max: 2, want: ""},
	}
	for _, tt := range tests {
		got := Truncate(tt.in, tt.max)
		if got != tt.want {
			t.Fatalf("Truncate(%q, %d) = %q, want %q", tt.in, tt.max, got, tt.want)
		}
		if len(got) > tt.max {
			t.Fatalf("Truncate(%q, %d) is %d bytes", tt.in, tt.max, len(got))
		}
	}
}

func BenchmarkTruncateLong(b *testing.B) {
	s := strings.Repeat("log line ", 100_000)
	for b.Loop() {
		_ = Truncate(s, 4096)
	}
}
