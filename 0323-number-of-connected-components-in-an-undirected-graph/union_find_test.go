package main

import "testing"

func TestCountComponents_union_find(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{0, 2},
				},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{2, 3},
					{4, 5},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponents_union_find(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("CountComponents_union_find() = %v, want %v", got, tt.want)
			}
		})
	}
}
