package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				strs: []string{"hello", "world", "i", "am", "apriil"},
			},
			want: "5#hello5#world1#i2#am6#apriil",
		},
		{
			name: "case 2",
			args: args{
				strs: []string{"today", "is", "a", "rainy", "day"},
			},
			want: "5#today2#is1#a5#rainy3#day",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encode(tt.args.strs)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				str: "5#hello5#world1#i2#am6#apriil",
			},
			want: []string{"hello", "world", "i", "am", "apriil"},
		},
		{
			name: "case 2",
			args: args{
				str: "5#today2#is1#a5#rainy3#day",
			},
			want: []string{"today", "is", "a", "rainy", "day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decode(tt.args.str)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_encodeAndDecode(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				input: []string{"hello", "world", "i", "am", "apriil"},
			},
			want: []string{"hello", "world", "i", "am", "apriil"},
		},
		{
			name: "case 2",
			args: args{
				input: []string{"today", "is", "a", "rainy", "day"},
			},
			want: []string{"today", "is", "a", "rainy", "day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeAndDecode(tt.args.input)

			assert.Equal(t, tt.want, got)
		})
	}
}
