package safepath

import (
	"errors"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		root, rel, want string
		err             error
	}{
		{root: "/srv/data", rel: "a/b.txt", want: "/srv/data/a/b.txt"},
		{root: "/srv/data", rel: ".", want: "/srv/data"},
		{root: "/srv/data/", rel: "a/../b", want: "/srv/data/b"},
		{root: "/srv/data", rel: "../etc/passwd", err: ErrEscapesRoot},
		{root: "/srv/data", rel: "/etc/passwd", err: ErrEscapesRoot},
		// A sibling whose name starts with the root's name is still outside.
		{root: "/srv/data", rel: "../data-old/secret", err: ErrEscapesRoot},
		{root: "/srv/data", rel: "a/../../data2", err: ErrEscapesRoot},
		// A name that merely starts with dots stays inside.
		{root: "/srv/data", rel: "..cache/x", want: "/srv/data/..cache/x"},
		{root: "relative/root", rel: "x", want: "relative/root/x"},
	}
	for _, tt := range tests {
		got, err := Join(tt.root, tt.rel)
		if !errors.Is(err, tt.err) {
			t.Fatalf("Join(%q, %q) error = %v, want %v", tt.root, tt.rel, err, tt.err)
		}
		if got != tt.want {
			t.Fatalf("Join(%q, %q) = %q, want %q", tt.root, tt.rel, got, tt.want)
		}
	}
}
