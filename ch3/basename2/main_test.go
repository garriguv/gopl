package main

import "testing"

func Test_basename(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a => a", args{"a"}, "a"},
		{"a.go => a", args{"a.go"}, "a"},
		{"a/b/c.go => c", args{"a/b/c.go"}, "c"},
		{"a/b.c.go => b.c", args{"a/b.c.go"}, "b.c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basename(tt.args.s); got != tt.want {
				t.Errorf("basename() = %v, want %v", got, tt.want)
			}
		})
	}
}
