package main

import "sync"

type DoubleList struct {
	head *ListNode //指向链表头部
	tail *ListNode //指向链表尾部
	len  int
	lock sync.Mutex
}

//列表的节点
type ListNode struct {
	pre   *ListNode //前驱节点
	next  *ListNode //后驱节点
	value interface{}
}

//从头部开始某个位置插入新节点
func (list *DoubleList) AddNodeFromHead(n int, v interface{}) { //输入一个要插入节点的位置和值
	list.lock.Lock()
	defer list.lock.Unlock()

	//如果索引超过此刻列表长度，直接报错
	if n != 0 && n >= list.len {
		panic("index out")

	}
	//先找到头部
	node := list.head
	//从头部往后遍历拿到n+1个位置的元素
	for i := 0; i <= n; i++ {
		node = node.next

	} //出循环就说明此时node就是n+1的位置

	//新节点
	newNode := new(ListNode)
	newNode.value = v

	//如果定位到的节点是空，表示整个列表是空的，将新节点设置为头部和尾部
	if node == nil {
		list.head = newNode
		list.tail = newNode

	} else { //如果定位到的节点不是空，那么他的前驱
		pre := node.pre
		//如果定位到的节点前驱是nil，就说明定位到的是头部节点，需要换头部
		if pre == nil {
			//将新节点加入老节点之前
			newNode.next = node //新节点的后驱指向原头部
			node.pre = newNode  //前驱指向新节点
			list.head = newNode //head指向新节点

		} else {
			//前驱节点不为空的话，就说明在列表中间，将定位到的节点的pre指向新节点，新节点的next指向
			pre.next = newNode
			newNode.pre = pre //新节点和定位的上一个节点相互连接

			newNode.next = node
			node.pre = newNode

		}

	}
	list.len++
}

//从头开始移除并返回某个位置的节点
func (list *DoubleList) PopFromeHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n >= list.len {
		return nil
	}
	node := list.head //获取头部
	//遍历拿到要移除的元素的位置
	for i := 1; i <= n; i++ {
		node = node.next

	}
	//要移除节点的前驱和后驱
	pre := node.pre
	next := node.next
	//如果前驱后驱都为零，说明要移除的节点是链表的唯一节点
	if pre == nil && next == nil {
		list.head = nil
		list.tail = nil
	} else if pre == nil {
		//表示要删除的是头部节点
		list.head = next
		next.pre = nil
	} else if next == nil {
		list.tail = pre
		pre.next = nil

	} else {
		pre.next = next
		next.pre = pre

	}
	list.len--
	return node
}
func main() {
	list := new(DoubleList)
	list.AddNodeFromHead(0, "我")
	list.AddNodeFromHead(0, "爱")
	list.AddNodeFromHead(0, "你")

}
