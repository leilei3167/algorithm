package main

import (
	"fmt"
	"sync"
)

type (
	Stack struct {
		top  *Node
		size int
		lock *sync.RWMutex
	}
	Node struct {
		value interface{}
		prev  *Node
	}
)

//创建一个结构体
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}

}

//入栈
func (this *Stack) Push(v interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &Node{v, this.top}
	this.top = n
	this.size++
}

//出栈
func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	//要取的元素就是top的元素
	if this.top == nil {
		return "没有元素了"

	} else {
		v := this.top.value
		this.top = this.top.prev //指向下一个
		this.size--
		return v
	}

}

//显示栈顶元素
func (this *Stack) list() interface{} {
	if this.top == nil {
		return "没有元素了"

	} else {
		v := this.top.value

		return v
	}

}

func main() {
	a := NewStack()
	a.Push("mother")
	a.Push("mfucker")
	a.Push(5.5555)
	a.Pop()
	fmt.Println(a.list())

}
