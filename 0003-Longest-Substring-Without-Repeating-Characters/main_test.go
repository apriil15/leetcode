package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			got := lengthOfLongestSubstring(tt.args.s)

			assert.Equal(t, got, tt.want, got)
		})
	}
}
