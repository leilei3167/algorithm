package main

import "fmt"

type (
	Stack struct {
		top  *node
		size int
	}
	node struct {
		value interface{}
		pre   *node
	}
)

//入栈
func (this *Stack) Push(v interface{}) {

	n := &node{v, this.top}
	this.top = n
	this.size++

}

//出栈
func (this *Stack) Pop() interface{} {
	if this.top == nil {
		panic("栈已空")

	}
	n := this.top
	this.top = n.pre //接到上一个去
	this.size--
	return n.value

}

//显示栈顶
func (this *Stack) Topis() interface{} {
	if this.top == nil {
		panic("栈已空")

	}
	n := this.top

	return n.value

}

//创建一个新的栈
func Newstack() *Stack {
	return &Stack{nil, 0}

}
func main() {
	stack1 := Newstack()
	stack1.Push("dhai")
	stack1.Push("123")
	stack1.Push(1233)
	stack1.Pop()
	fmt.Println(stack1.Topis())
}
