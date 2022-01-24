package main

import (
	"fmt"
	"sync"
)

/* 广度遍历法遍历二叉树，使用队列进行辅助 */
//构造二叉树
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

//构造队列，链表的节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

//链表队列
type LinkQueue struct {
	root *LinkNode //起点
	size int
	lock sync.Mutex
}

//入队
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//如果队列为空，则增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		//否则新元素插入末尾
		newNode := new(LinkNode)
		newNode.Value = v
		//遍历
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}
	queue.size++
}

//出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}
	//顶部先出
	topNode := queue.root
	v := topNode.Value
	queue.root = topNode.Next
	queue.size--
	return v

}

//队元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}

//对二叉树的遍历，放入根节点
func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return

	}
	//新建队列
	queue := new(LinkQueue)
	//根节点放入
	queue.Add(treeNode)

	for queue.size > 0 {
		//不断出队列
		element := queue.Remove() //第一轮先取出a打印，c入队后循环到第二轮，此刻取出的是上一轮先入的b
		//打印节点值
		fmt.Println(element.Data)
		//左儿子非空，入队列
		if element.Left != nil {
			queue.Add(element.Left) //b先入队
		}
		//右儿子非空，入队列
		if element.Right != nil {
			queue.Add(element.Right) //c后入队
		}

	}

}
func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "b"}
	t.Right = &TreeNode{Data: "c"}
	t.Left.Left = &TreeNode{Data: "d"}
	t.Left.Right = &TreeNode{Data: "e"}
	t.Right.Left = &TreeNode{Data: "f"}
	LayerOrder(t)
}
