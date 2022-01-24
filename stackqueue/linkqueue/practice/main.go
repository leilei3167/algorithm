package main

import "fmt"

/* 队列是先进先出，主要区别是入队要入到最后一个 */
type head struct {
	root *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

//入队
func (queue *head) getin(v interface{}) {
	if queue.root == nil {
		queue.root = new(node)
		queue.root.value = v

	} else {
		//要把元素放到链表末尾去
		newnode := new(node)
		newnode.value = v

		topnode := queue.root
		for topnode.next != nil {
			topnode = topnode.next

		} //出循环就说明已经到尾部

		topnode.next = newnode //将新元素链接到其尾部

	}
	queue.size = queue.size + 1

}

//出队
func (queue *head) out() interface{} {
	if queue.root == nil {
		return "没有可取的"
	} else {
		//要出的就是第一个元素
		v := queue.root.value
		queue.root = queue.root.next
		queue.size = queue.size - 1
		return v
	}

}

//显示下一个要出的元素
func (queue *head) list() interface{} {
	if queue.root == nil {
		return "没有元素"
	} else {
		//要出的就是第一个元素

		return queue.root.value
	}

}

func main() {
	//测试
	a := new(head)
	a.getin(56)
	a.getin(59)
	a.getin("ni ni ni")
	fmt.Println(a.list())
}
