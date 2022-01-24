package main

//队列用链表实现
type head struct {
	root *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

//入队
func (a *head) push(v interface{}) {
	if a.root == nil {
		a.root = new(node)
		a.root.value = v
	} else {
		newnode := new(node)
		newnode.value = v //创建新节点 然后接到末尾去
		//如何找到末尾？
		//此时第一个元素就是a.root指向的
		firstnode := a.root
		for firstnode.next != nil {
			firstnode = firstnode.next

		} //出循环就说明到末尾
		firstnode.next = newnode

	}
	a.size = a.size + 1

}

//出队
func (a *head) out() interface{} {
	if a.root == nil {
		return "无元素"

	} else {
		v := a.root.value
		a.root = a.root.next
		a.size = a.size - 1
		return v

	}

}
