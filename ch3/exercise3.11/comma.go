package comma

import (
	"bytes"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// check if the number has a sign
	var sign string
	if neg := strings.Index(s, "-"); neg >= 0 {
		s, sign = s[neg+1:], s[:neg+1]
	}

	// only insert commas before the period in floating point numbers.
	var r string
	if period := strings.Index(s, "."); period >= 0 {
		s, r = s[:period], s[period:]
	}

	// early return for small numbers
	n := len(s)
	if n <= 3 {
		return sign + s + r
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
	return sign + buf.String()
}
