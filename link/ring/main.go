package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func add(head *CatNode, newcat *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil { //是的话自己就是作为第一只猫来添加
		head.no = newcat.no
		head.name = newcat.name
		head.next = head //形成环状，注意此处不能指向newcat
		fmt.Println("这只猫已经加入环形链表")
		return
	}
	//如果加的不是第一只
	temp := head //临时变量temp帮忙找到环形最后一个节点
	for {
		if temp.next == head {
			break
		}
		temp = temp.next

	}

	temp.next = newcat
	newcat.next = head

}

//输出链表
func list(head *CatNode) {
	fmt.Println("环形链表的具体情况：")
	temp := head
	if temp.next == nil {
		fmt.Println("目前是空链表！！！")
	}
	for {
		fmt.Printf("[%d,%s]--->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next

	}
}

//删除猫
/* 思路：
1.先让temp指向head
2.新建一个helper，指向最后一个元素
3.让temp和要删除的id比较，则通过helper删除（必须考虑要删除的就是头节点怎么办？

*/
func del(head *CatNode, no int) {
	temp := head
	helper := head
	if temp.next == nil { //先看是否是空链表
		fmt.Println("这是空的无法删除")
		return
	}
	//如果只有一个节点
	if temp.next == head { //只有一个节点
		temp.next = nil
		return
	}
	//两个以上节点的话
	for {
		if temp.next == head { //说明到最后一个
			break
		}
		if temp.no == no {
			//找到，也可以在此删除
			helper.next = temp.next
			break
		}
		temp = temp.next
		helper = helper.next //一旦找到要删除的节点，就需要helper来删除
	}

}

func main() {
	//初始化一个环形链表的头节点
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}

	cat3 := &CatNode{
		no:   3,
		name: "lol",
	}

	add(head, cat1)
	add(head, cat2)
	add(head, cat3)
	list(head)
}
