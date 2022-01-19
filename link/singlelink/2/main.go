package main

import "fmt"

type HeroNode struct {
	No       int
	Name     string
	Nickname string
	next     *HeroNode
}

//增删改查单链表
//第一种插入方式，在单链表的最后加入
func insert(head *HeroNode, newhero *HeroNode) {
	//1.先找到该链表最后一个节点
	//2.创建一个辅助节点，
	temp := head
	for {
		if temp.next == nil { //表示找到了最后一个，因为最后一个节点指向的是空
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	} //跳出这个循环的时候就说明一定是到了链表末尾
	//3.将newhero加入链表的最后
	temp.next = newhero //让nil值指向一个新加入的节点

}

//第二种方式，要求根据序号来进行插入排序
/* 1.找到适当的节点位置 */
func insert1(head *HeroNode, newhero *HeroNode) {
	temp := head
	flag := true
	//让插入的节点的序号和temp的下一个节点的序号进行比较
	for {
		if temp.next == nil { //说明到链表最后
			break

		} else if temp.next.No > newhero.No {
			//说明newhero就应该插入到此时temp的后面
			break

		} else if temp.next.No == newhero.No {
			flag = false
			break
		}
		temp = temp.next //往下走

	}
	if !flag {
		fmt.Println("已存在相同的id!!")
		return
	} else {
		newhero.next = temp.next //要插入的节点，必须先指向上一个节点指向的下一个节点
		temp.next = newhero      //之后再将

	}
}

//删除一个节点
func del(head *HeroNode, id int) {

	temp := head
	flag := false
	//找到要删除的节点
	for {
		if temp.next == nil {
			fmt.Println("找不到对应id的内容！！")
			return
		} else if temp.next.No == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //找到了的话就
		temp.next = temp.next.next //跳过要被删除的节点
	}

}

//显示链表的所有信息
func list(head *HeroNode) {
	temp := head
	//先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("空链表无法显示")
		return
	}
	//遍历
	for {
		fmt.Printf("[%d,%s,%s]--->", temp.next.No, temp.next.Name, temp.next.Nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	//创建一个头结点
	head := &HeroNode{}

	//创建新Heronode
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		Nickname: "及时雨",
	}

	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		Nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		Name:     "鲁智深",
		Nickname: "和尚",
	}
	/* hero4 := &HeroNode{
		No:       3,
		Name:     "吴用",
		Nickname: "智多星",
	} */
	//加入
	insert1(head, hero3)
	insert1(head, hero1)
	insert1(head, hero2)
	/* insert1(head, hero4) */

	list(head)
	fmt.Println()
	//删除
	del(head, 1)
	

	list(head)

}
