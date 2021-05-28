package main

import "testing"

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	datas := []struct {
		name string
		args args
		want int
	}{
		{name: "first", args: args{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}}, want: 16},
		{name: "second", args: args{[][]int{{1}, {0}}}, want: 4},
	}
	for _, data := range datas {
		t.Run(data.name, func(t *testing.T) {
			if got := islandPerimeter(data.args.grid); got != data.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, data.want)
			}
		})
	}
}
