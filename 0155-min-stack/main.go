package main

func main() {

}

type MinStack struct {
	stack []Node
}

type Node struct {
	value  int
	curMin int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	var curMin int
	if len(this.stack) == 0 {
		curMin = val
	} else {
		curMin = min(val, this.GetMin())
	}

	this.stack = append(this.stack, Node{
		value:  val,
		curMin: curMin,
	})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].curMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
