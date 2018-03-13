package main

import "testing"

func TestPopCount(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"127", args{127}, 7},
		{"128", args{128}, 1},
		{"18446744073709551614", args{18446744073709551614}, 63},
		{"18446744073709551615", args{18446744073709551615}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args.x); got != tt.want {
				t.Errorf("PopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopCountShift(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"127", args{127}, 7},
		{"128", args{128}, 1},
		{"18446744073709551614", args{18446744073709551614}, 63},
		{"18446744073709551615", args{18446744073709551615}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCountShift(tt.args.x); got != tt.want {
				t.Errorf("PopCountShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(18446744073709551615)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(18446744073709551615)
	}
}
