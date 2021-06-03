package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{"()"}, want: true},
		{name: "Example 2", args: args{"()[]{}"}, want: true},
		{name: "Example 3", args: args{"(]"}, want: false},
		{name: "Example 4", args: args{"([)]"}, want: false},
		{name: "Example 5", args: args{"{[]}"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
