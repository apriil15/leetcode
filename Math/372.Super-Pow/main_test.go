package main

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{2, []int{3}}, want: 8},
		{name: "2", args: args{2, []int{1, 0}}, want: 1024},
		{name: "3", args: args{1, []int{4, 3, 3, 8, 5, 2}}, want: 1},
		{name: "4", args: args{2147483647, []int{2, 0, 0}}, want: 1198},
		{name: "5", args: args{1337 * 2, []int{8, 7}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
