// Package safepath joins untrusted relative paths onto a trusted root.
package safepath

import (
	"errors"
	"path/filepath"
	"strings"
)

// ErrEscapesRoot reports a path that would resolve outside the root.
var ErrEscapesRoot = errors.New("safepath: path escapes root")

// Join returns root joined with rel. It fails when rel is absolute or when
// the cleaned result is not root itself or inside it.
func Join(root, rel string) (string, error) {
	if filepath.IsAbs(rel) {
		return "", ErrEscapesRoot
	}
	root = filepath.Clean(root)
	joined := filepath.Join(root, rel)
	within, err := filepath.Rel(root, joined)
	if err != nil {
		return "", ErrEscapesRoot
	}
	if within == ".." || strings.HasPrefix(within, ".."+string(filepath.Separator)) {
		return "", ErrEscapesRoot
	}
	return joined, nil
}
