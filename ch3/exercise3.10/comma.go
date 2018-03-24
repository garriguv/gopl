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
	for len(s) > 3 {
		comma := len(s) % 3
		if comma == 0 {
			comma = 3
		}
		buf.WriteString(s[:comma])
		buf.WriteByte(',')
		s = s[comma:]
	}
	buf.WriteString(s)
	buf.WriteString(r)
	return buf.String()
}
