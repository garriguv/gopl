package tempconv

import "testing"

func TestCToF(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{"absolute", args{AbsoluteZeroC}, -459.66999999999996},
		{"freezing", args{FreezingC}, 32},
		{"boiling", args{BoilingC}, 212},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToF(tt.args.c); got != tt.want {
				t.Errorf("CToF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCToK(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{"absolute", args{AbsoluteZeroC}, 0},
		{"freezing", args{FreezingC}, 273.15},
		{"boiling", args{BoilingC}, 373.15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.args.c); got != tt.want {
				t.Errorf("CToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{"absolute", args{-459.66999999999996}, AbsoluteZeroC},
		{"freezing", args{32}, FreezingC},
		{"boiling", args{212}, BoilingC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToC(tt.args.f); got != tt.want {
				t.Errorf("FToC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFToK(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{"absolute", args{-459.66999999999996}, 0},
		{"freezing", args{32}, 273.15},
		{"boiling", args{212}, 373.15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToK(tt.args.f); got != tt.want {
				t.Errorf("FToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKToC(t *testing.T) {
	type args struct {
		k Kelvin
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{"absolute", args{0}, AbsoluteZeroC},
		{"freezing", args{273.15}, FreezingC},
		{"boiling", args{373.15}, BoilingC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToC(tt.args.k); got != tt.want {
				t.Errorf("KToC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKToF(t *testing.T) {
	type args struct {
		k Kelvin
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{"absolute", args{0}, -459.66999999999996},
		{"freezing", args{273.15}, 32},
		{"boiling", args{373.15}, 212},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToF(tt.args.k); got != tt.want {
				t.Errorf("KToF() = %v, want %v", got, tt.want)
			}
		})
	}
}
