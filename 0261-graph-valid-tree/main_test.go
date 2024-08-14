package main

import (
	"testing"
)

func TestValidTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidTree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("ValidTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
