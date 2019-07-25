package count

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestDifferent(t *testing.T) {
	tests := []struct {
		h1   [32]byte
		h2   [32]byte
		want int
	}{
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 125},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := Different(tt.h1, tt.h2)

			if want, got := tt.want, got; want != got {
				t.Errorf("unexpected output: want %v got %v", want, got)
			}
		})
	}
}

func TestDifferentSlice(t *testing.T) {
	tests := []struct {
		h1   []byte
		h2   []byte
		want int
	}{
		{[]byte{2}, []byte{7}, 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := differentSlice(tt.h1, tt.h2)

			if want, got := tt.want, got; want != got {
				t.Errorf("unexpected output: want %v got %v", want, got)
			}
		})
	}
}
