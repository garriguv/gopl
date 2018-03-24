package exercise3_12

import "strings"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		if i := strings.IndexRune(s2, r); i >= 0 {
			s2 = s2[:i] + s2[i+1:]
		} else {
			return false
		}
	}
	return true
}
