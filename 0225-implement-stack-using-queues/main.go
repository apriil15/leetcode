package main

func main() {

}

type MyStack struct {
	s []int
}

func Constructor() MyStack {
	// 100: At most 100 calls
	return MyStack{s: make([]int, 0, 100)}
}

func (this *MyStack) Push(x int) {
	this.s = append(this.s, x)
}

// time: O(n)
//
// move first to last
//
//	1, 2, 3, 4
//	   2, 3, 4, 1
//	      3, 4, 1, 2
//
// remove first
//
//	4, 1, 2, 3
//	   1, 2, 3
func (this *MyStack) Pop() int {
	// move first to last n-1 times
	for i := 0; i < len(this.s)-1; i++ {
		first := this.s[0]
		this.s = append(this.s[1:], first)
	}

	// still one time
	result := this.s[0]
	this.s = this.s[1:]

	return result
}

func (this *MyStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.s) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
