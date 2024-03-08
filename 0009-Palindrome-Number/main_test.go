package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{121}, want: true},
		{name: "2", args: args{-121}, want: false},
		{name: "3", args: args{10}, want: false},
		{name: "4", args: args{-101}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
