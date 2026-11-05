package textutil

import (
	"testing"
	"unicode/utf8"
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
		{in: "héllo wörld", max: 9, want: "héllo…"},
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
		if !utf8.ValidString(got) {
			t.Fatalf("Truncate(%q, %d) = %q is not valid UTF-8", tt.in, tt.max, got)
		}
	}
}
