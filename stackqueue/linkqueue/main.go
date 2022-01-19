package main

import (
	"fmt"
	"sync"
)

/* 队列先进先出，和栈操作顺序相反，我们这里只实现入队，和出队操作，其他操作和栈一样。
因为先进先出，所以要遍历链表确认是放在最后 */
//创建链表头部和节点
type linkqueue struct {
	root *linknode
	size int
	lock sync.Mutex
}

//节点
type linknode struct {
	next  *linknode
	value string
}

//入队
func (queue *linkqueue) in(v string) {
	//判断是否已有元素
	if queue.root == nil { //没有元素的话就创建成为新的节点
		queue.root = new(linknode)
		queue.root.value = v

	} else { //就把元素插在末尾
		//新节点
		newNode := new(linknode)
		newNode.value = v
		//遍历到链表尾部
		nownode := queue.root //从头部开始
		for nownode.next != nil {
			nownode = nownode.next

		} //直到下一个地址为nil就说明到尾部，此刻nownode就是最后一个元素
		nownode.next = newNode //把新节点放到尾部

	}

	queue.size = queue.size + 1
}

//出队
func (queue *linkqueue) out() string {
	if queue.root == nil {
		return "么有元素了！"
	}
	topnode := queue.root
	v := topnode.value
	// 将顶部元素的后继链接链上
	queue.root = topnode.next

	queue.size = queue.size - 1
	return v

}

//显示下一个可以出队的值
func (queue *linkqueue) next() string {
	if queue.root == nil {
		return "没有元素了"

	}
	return queue.root.value

}
func main() {
	//测试
	a := new(linkqueue)
	a.in("雷磊1")
	a.in("郝运2")
	a.in("杨真3")
	fmt.Println("下一个出队的元素是", a.next())

}
