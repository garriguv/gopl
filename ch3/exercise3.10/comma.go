package comma

import (
	"bytes"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// only insert commas before the period in floating point numbers.
	var r string
	if period := strings.Index(s, "."); period >= 0 {
		s, r = s[:period], s[period:]
	}

	// early return for small numbers
	n := len(s)
	if n <= 3 {
		return s + r
	}

	var buf bytes.Buffer
	if n%3 != 0 {
		buf.WriteString(s[:n%3])
		buf.WriteByte(',')
		s = s[n%3:]
	}
	for len(s) > 3 {
		buf.WriteString(s[:3])
		buf.WriteByte(',')
		s = s[3:]
	}
	buf.WriteString(s)
	buf.WriteString(r)
	return buf.String()
}
