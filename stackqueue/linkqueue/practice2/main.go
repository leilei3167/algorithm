package main

import "fmt"

/* 链表队列 */
type (
	Queue struct {
		top  *node
		size int
	}
	node struct {
		value interface{}
		next  *node
	}
)

//创建结构体
func NewQueue() *Queue {
	return &Queue{nil, 0}
}

//入队
func (this *Queue) Push(v interface{}) {
	//判断是否已有元素，没有的话就要把头指向第一个新节点，后续头指向不再改变
	//要出队直接就是出top指向的元素
	if this.top == nil {
		newnode := new(node)
		newnode.value = v
		this.top = newnode

	} else {
		newnode := new(node)
		newnode.value = v
		//遍历到尾部
		nownode := this.top //从头开始
		for nownode.next != nil {
			nownode = nownode.next

		}
		nownode.next = newnode
	}
	this.size++

}

//显示元素
func (this *Queue) Next() interface{} {
	if this.size == 0 {
		panic("没有元素")
	}
	//要出的元素就是top指向的
	v := this.top.value

	return v

}

//出队
func (this *Queue) Pop() interface{} {
	if this.size == 0 {
		panic("没有元素")
	}
	//要出的元素就是top指向的
	v := this.top.value

	this.top = this.top.next //接到下一个
	this.size--
	return v

}
func main() {
	queue1 := NewQueue()
	queue1.Push(12)
	queue1.Push(4325252)
	queue1.Push(313213)
	queue1.Push(5435436)
	queue1.Push(54354)
	queue1.Push(1767)

	fmt.Println(queue1.Next())
}
