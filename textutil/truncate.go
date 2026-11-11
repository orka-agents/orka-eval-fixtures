// Package textutil holds small text helpers.
package textutil

import "unicode/utf8"

// Ellipsis is appended to text that Truncate shortens.
const Ellipsis = "…"

// TruncateRunes returns s unchanged when it has at most maxRunes characters.
// Otherwise it returns its first maxRunes-1 characters followed by Ellipsis.
// A maxRunes below 1 returns "".
func TruncateRunes(s string, maxRunes int) string {
	if maxRunes < 1 {
		return ""
	}
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}
	count := 0
	for i := range s {
		if count == maxRunes-1 {
			return s[:i] + Ellipsis
		}
		count++
	}
	return s
}

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
