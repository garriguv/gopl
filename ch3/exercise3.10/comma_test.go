package comma

import "testing"

func Test_comma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"123 => 123", args{"123"}, "123"},
		{"1234 => 1,234", args{"1234"}, "1,234"},
		{"123456789 => 123,456,789", args{"123456789"}, "123,456,789"},
		{"123.4 => 123.4", args{"123.4"}, "123.4"},
		{"1234.5 => 1,234.5", args{"1234.5"}, "1,234.5"},
		{"123456789.123 => 123,456,789.123", args{"123456789.123"}, "123,456,789.123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args.s); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma("1234567.123")
	}
}
