package main

import "testing"

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first",
			args: args{
				employees: []*Employee{
					{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
					{Id: 2, Importance: 3, Subordinates: []int{}},
					{Id: 3, Importance: 3, Subordinates: []int{}}},
				id: 1},
			want: 11,
		},
		{name: "second",
			args: args{
				employees: []*Employee{
					{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
					{Id: 2, Importance: 3, Subordinates: []int{4}},
					{Id: 3, Importance: 4, Subordinates: []int{}},
					{Id: 4, Importance: 1, Subordinates: []int{}}},
				id: 1},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
