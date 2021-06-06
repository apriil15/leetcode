package main

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "first", args: args{a: "11", b: "1"}, want: "100"},
		{name: "second", args: args{a: "1010", b: "1011"}, want: "10101"},
		{name: "third", args: args{a: "1", b: "111"}, want: "1000"},
		{name: "fourth", args: args{a: "110", b: "1"}, want: "111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
