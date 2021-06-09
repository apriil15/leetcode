package main

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "first", args: args{s: "abcdefg", k: 2}, want: "bacdfeg"},
		{name: "second", args: args{s: "abcd", k: 2}, want: "bacd"},
		{name: "third", args: args{s: "abcdefgh", k: 3}, want: "cbadefhg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
