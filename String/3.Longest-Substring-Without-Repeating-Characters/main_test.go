package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first", args: args{s: "abcabcbb"}, want: 3},
		{name: "second", args: args{s: "bbbbb"}, want: 1},
		{name: "third", args: args{s: "pwwkew"}, want: 3},
		{name: "fourth", args: args{s: ""}, want: 0},
		{name: "fifth", args: args{s: "aab"}, want: 2},
		{name: "6", args: args{s: "dvdf"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
