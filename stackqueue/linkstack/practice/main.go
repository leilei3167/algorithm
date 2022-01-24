package main

import "fmt"

/* 链表实现栈结构，思路：
1.创建起点，和节点
2.入栈时，首先考虑是否为第一个元素，如果是的话，直接创建新节点，并将起点的next指向这个新创建的节点
如果不是第一个，那就要插入到第一个，将要压入的元素的next指向原第一个元素，并将起点的next指向要插入的元素
并且栈空间要加一
3，出栈时，先考虑有没有元素，如果没有，则要返回错误。如果有，要返回的值就为现在第一个元素，即起点指向的元素的值，
要将要将起点指向重新修改到第二个元素
*/
//创建起点和节点

type start struct {
	root *node //指向节点
	size int   //栈空间的大小
}

//节点
type node struct {
	value interface{} //可存入任意类型的值
	next  *node
}

//入栈的方法
func (stack *start) push(v interface{}) {
	//先判断是否是第一个元素
	if stack.root == nil {
		stack.root = new(node) //是第一个元素就创建一个
		stack.root.value = v

	} else { //如果已有节点
		prenode := stack.root //原来的第一个元素
		//创建的要插入的元素
		newnode := new(node)
		newnode.value = v
		//将新元素插入
		newnode.next = prenode
		stack.root = newnode

	}
	stack.size = stack.size + 1

}

//出栈
func (stack *start) out() interface{} {
	if stack.root == nil {
		return "没有元素可供取出！！"

	} else { //要取出的值就是起点指向的值

		v := stack.root.value
		stack.root = stack.root.next
		stack.size = stack.size - 1
		return v

	}

}

//显示下一个要出栈的元素
func (stack *start) list() interface{} {
	if stack.root == nil {
		return "没有元素可供取出！！"

	} else { //要取出的值就是起点指向的值

		return stack.root.value

	}

}

func main() {
	//test
	a := new(start)

	a.push("德玛西亚")
	a.push(5.6)
	a.push("china no.1!")
	a.out()
	fmt.Println(a.list())

}
