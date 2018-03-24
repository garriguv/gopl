package exercise3_12

import "testing"

func TestAnagram(t *testing.T) {
	type args struct {
		s1, s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{" => true", args{"", ""}, true},
		{"123, 321 => true", args{"123", "321"}, true},
		{"133, 321 => false", args{"133", "321"}, false},
		{"1234, 321 => false", args{"1234", "321"}, false},
		{"rail safety, fairy tales => true", args{"rail safety", "fairy tales"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anagram("rail safety", "fairy tales")
	}
}
