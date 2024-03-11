package main

import (
	"testing"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "base", args: args{1}, want: "1"},
		{name: "first", args: args{2}, want: "11"},
		{name: "second", args: args{3}, want: "21"},
		{name: "third", args: args{4}, want: "1211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_say(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "first", args: args{"1"}, want: "11"},
		{name: "second", args: args{"11"}, want: "21"},
		{name: "third", args: args{"21"}, want: "1211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := say(tt.args.str); got != tt.want {
				t.Errorf("say() = %v, want %v", got, tt.want)
			}
		})
	}
}
