package main

import (
	"testing"
)

func Test_findLHS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first", args: args{[]int{1, 3, 2, 2, 5, 2, 3, 7}}, want: 5},
		{name: "second", args: args{[]int{1, 2, 3, 4}}, want: 2},
		{name: "third", args: args{[]int{1, 1, 1, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
