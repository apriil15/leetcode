package main

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}}},
		{name: "2", args: args{head: nil}, want: nil},
		{name: "3", args: args{head: &ListNode{1, nil}}, want: &ListNode{1, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
