package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "first", args: args{[]string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "second", args: args{[]string{"dog", "racecar", "car"}}, want: ""},
		{name: "third", args: args{[]string{"a"}}, want: "a"},
		{name: "fourth", args: args{[]string{"ab", "a"}}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
