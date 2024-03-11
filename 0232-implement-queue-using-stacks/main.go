package main

func main() {

}

// 2 stacks
// one is for push, the other is for pop
//
// -   -   -   -
// - 4 -   - 1 -
// - 3 -   - 2 -
// - 2 -   - 3 -
// - 1 -   - 4 -
// - - -   - - -
type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: []int{},
		popStack:  []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

// time: O(1)
func (this *MyQueue) Pop() int {
	// only when popStack is empty, time: O(n)
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			last := this.pushStack[len(this.pushStack)-1]
			this.popStack = append(this.popStack, last)

			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}

	last := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return last
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) > 0 {
		return this.popStack[len(this.popStack)-1]
	}

	return this.pushStack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack)+len(this.popStack) == 0
}
