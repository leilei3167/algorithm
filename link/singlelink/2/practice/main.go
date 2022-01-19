package main

import "fmt"

/* 实现单链表的增删改查 */
//创建结构体
type Student struct {
	name  string
	no    int
	score int
	next  *Student
}

//增加节点:方法一：新节点直接放在链表末尾
func add1(head *Student, newstu *Student) {
	//1.先找到最后一个节点的位置
	temp := head //head是一个地址
	for {        //一直循环到找到末尾为止
		if temp.next == nil { //指向nil就说明已经到了末尾
			break
		}
		temp = temp.next //将temp指向的位置不断后移
	}
	//找到末尾后，添加
	temp.next = newstu

}

//增加节点:方法二：新节点按照序号顺序插入
func add2(head *Student, newstu *Student) {
	flag := true
	temp := head
	//让要插入的节点的序号和下一个进行比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newstu.no { //判断下一个节点的no是否大于新加入的
			break
		} else if temp.next.no == newstu.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在相同id！！")

	} else {
		newstu.next = temp.next
		temp.next = newstu //这一步实现插入
	}
}

//显示节点
func list(head *Student) {
	//先判断是否是一个空链表
	temp := head
	if temp.next == nil {
		fmt.Println("这是个空链表，还没有任何数据！！")
		return
	}
	for {
		fmt.Printf("[%s,%d,%d]-->", temp.next.name, temp.next.no, temp.next.score)
		temp = temp.next
		if temp.next == nil {
			fmt.Println()
			break
		}

	}

}

//删除节点
func del(head *Student, n int) {

	temp := head
	for {
		if temp.next.no == n {
			break
		} else if temp.next == nil {
			fmt.Println("没有对应的iD！！！")
			return
		}

		temp = temp.next

	}
	temp.next = temp.next.next

}

func main() {
	//头结点
	head := &Student{}
	head2 := &Student{}
	//要加入的节点
	stu1 := &Student{
		name:  "jace",
		no:    1,
		score: 18,
	}

	stu2 := &Student{
		name:  "bob",
		no:    2,
		score: 145,
	}

	stu3 := &Student{
		name:  "juli",
		no:    6,
		score: 19,
	}

	stu4 := &Student{
		name:  "kevien",
		no:    5,
		score: 100,
	}
	fmt.Println("----------第一种方式增加-----------")
	fmt.Println("最开始head的next：", head.next) //开始head中的next地址是nil
	add1(head, stu1)
	fmt.Println("增加了stu1后的head的next：", head.next) //增加后head中的next就是stu1的地址，&{jace 1 18 <nil>}
	add1(head, stu2)
	fmt.Println("stu1的next", stu1.next) //增加stu2之后stu1的地址就是stu2的地址
	add1(head, stu3)
	add1(head, stu4)

	list(head)
	fmt.Println("----------第二种方式增加-----------")
	add2(head2, stu4)
	add2(head2, stu3)
	add2(head2, stu1)
	add2(head2, stu2)
	list(head2)
	fmt.Println("----------删除id为1和5的数-----------")
	del(head2, 1)
	del(head2, 5)
	list(head2)

}
