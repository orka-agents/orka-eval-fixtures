// Package textutil holds small text helpers.
package textutil

import "unicode/utf8"

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
	// Every byte of an ASCII prefix is a whole character, so the prefix can
	// be cut at any byte. budget < len(s) also holds when maxBytes is so
	// negative that the subtraction above wrapped around.
	if budget < len(s) && isASCII(s[:budget]) {
		return s[:budget] + Ellipsis
	}
	end := 0
	for i, r := range s {
		size := utf8.RuneLen(r)
		if size < 0 {
			size = 1
		}
		if i+size > budget {
			break
		}
		end = i + size
	}
	return s[:end] + Ellipsis
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}
