package main

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 1", args: args{num1: "11", num2: "123"}, want: "134"},
		{name: "Example 2", args: args{num1: "456", num2: "77"}, want: "533"},
		{name: "Example 3", args: args{num1: "0", num2: "0"}, want: "0"},
		{name: "Example 4", args: args{num1: "87", num2: "78"}, want: "165"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
