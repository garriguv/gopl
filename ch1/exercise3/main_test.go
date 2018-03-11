package main

import (
	"strings"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	args := []string{"this", "is", "a", "test"}
	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 0; j < len(args); j++ {
			s += sep + args[j]
			sep = " "
		}
	}
}

func BenchmarkEcho3(b *testing.B) {
	args := []string{"this", "is", "a", "test"}
	for i := 0; i < b.N; i++ {
		_ = strings.Join(args, " ")
	}
}
