// Package textutil holds small text helpers.
package textutil

// Ellipsis is appended to text that Truncate shortens.
const Ellipsis = "…"

// Truncate returns s unchanged when it fits in maxBytes. Otherwise it returns
// the longest prefix of s that, followed by Ellipsis, fits in maxBytes. It
// never splits a UTF-8 sequence. A maxBytes too small for Ellipsis returns "".
func Truncate(s string, maxBytes int) string {
	if len(s) <= maxBytes {
		return s
	}
	budget := maxBytes - len(Ellipsis)
	if budget < 0 {
		return ""
	}
	// Slicing by byte offset avoids decoding every rune of long inputs.
	return s[:budget] + Ellipsis
}
