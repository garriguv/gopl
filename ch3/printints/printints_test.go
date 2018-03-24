package printints

import "testing"

func Test_intsToString(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"[]int{1, 2, 3} => [1, 2, 3]", args{[]int{1, 2, 3}}, "[1, 2, 3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intsToString(tt.args.values); got != tt.want {
				t.Errorf("intsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
