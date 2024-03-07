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

func (this *MyStack) Pop() int {
	result := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]

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
