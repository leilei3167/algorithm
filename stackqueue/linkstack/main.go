package main

import (
	"fmt"
	"sync"
)

/* 链表实现栈 */
//表头
type LinkStack struct {
	root *LinkNode  //起点
	size int        //栈元素数量
	lock sync.Mutex //并发安全
}

//节点
type LinkNode struct {
	next  *LinkNode
	value string
}

//入栈
func (stack *LinkStack) push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	//要先判断是否是第一个元素
	if stack.root == nil {
		//说明还没有元素，那么要增加新节点
		stack.root = new(LinkNode) //增加新节点 并与root链接
		stack.root.value = v

	} else {
		//说明已有元素，就要将新元素插入到第一个
		//原来的链表
		preNode := stack.root
		//创建新节点
		newNode := new(LinkNode)
		newNode.value = v
		//将原来的链表链接到新节点上
		newNode.next = preNode
		//将新节点放到头部
		stack.root = newNode

	}
	//栈元素+1
	stack.size = stack.size + 1

}

//出栈
func (stack *LinkStack) pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	//判断元素是否已空
	if stack.root == nil {
		panic("栈中没有元素！！")

	}
	//顶部元素先出栈
	topNode := stack.root
	v := topNode.value

	//后续元素与root链接
	stack.root = topNode.next
	stack.size = stack.size - 1
	return v
}

//显示栈顶元素
func (stack *LinkStack) top() string {
	if stack.root == nil {
		return "栈中已经没有元素了"

	}
	v := stack.root.value
	return v

}

// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	//测试
	a := new(LinkStack)
	a.push("雷磊1")
	a.push("郝运2")
	a.push("杨真3")
	fmt.Println("此时栈大小为", a.Size(), "栈顶元素为", a.top())
	//出栈
	a.pop()
	fmt.Println("此时栈大小为", a.Size(), "栈顶元素为", a.top())

}
