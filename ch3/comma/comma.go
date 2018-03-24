package comma

import "strings"

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// only insert commas before the period in floating point numbers.
	if period := strings.Index(s, "."); period >= 0 {
		return comma(s[:period]) + s[period:]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
