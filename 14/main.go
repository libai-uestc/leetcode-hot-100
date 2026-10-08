package main

type MinStack struct {
	stack [][2]int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([][2]int, 0),
	}
}

func (this *MinStack) Push(value int) {
	minVal := value
	if len(this.stack) > 0 {
		currentMin := this.stack[len(this.stack)-1][1]
		if currentMin < minVal {
			minVal = currentMin
		}
	}
	this.stack = append(this.stack, [2]int{value, minVal})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
