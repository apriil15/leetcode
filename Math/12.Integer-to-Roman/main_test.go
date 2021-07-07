package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{3}, want: "III"},
		{name: "2", args: args{4}, want: "IV"},
		{name: "3", args: args{9}, want: "IX"},
		{name: "4", args: args{58}, want: "LVIII"},
		{name: "5", args: args{1994}, want: "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
