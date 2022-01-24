package main

import "fmt"

/* 创建起点 和节点 */
type start struct {
	root *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

//入栈

func (stack *start) push(v interface{}) {
	//判断是否为第一个值
	if stack.root == nil {
		stack.root = new(node)
		stack.root.value = v

	} else {
		//如果不是第一个元素，那么就要插入
		prenode := stack.root
		newnode := new(node)
		newnode.value = v
		//链接
		newnode.next = prenode
		stack.root = newnode

	}
	stack.size = stack.size + 1

}

//出栈
func (stack *start) out() interface{} {
	if stack.root == nil {
		return "没有元素"

	} else {
		//要取的就是栈顶
		v := stack.root.value

		//重新链接
		stack.root = stack.root.next
		stack.size = stack.size - 1
		return v

	}

}

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
	a.out()
	fmt.Println(a.list())
}
