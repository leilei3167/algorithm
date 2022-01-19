package main

import "fmt"

type linklist struct {
	data     int
	nextnode *linklist
} //创建一个结构体，包含数据以及下一个节点的地址

func main() {
	//创建第一个一个节点
	node := new(linklist)
	node.data = 100                  //未初始化的将是默认的零值
	fmt.Printf("node的类型：%T\n", node) //因为new创建的node是一个指针类型，node的类型：*main.linklist

	node3 := linklist{data: 54}
	fmt.Printf("node3的类型：%T\n", node3) //直接声明定义的node3的类型：main.linklist

	//创建第二个节点
	node1 := new(linklist)
	node1.data = 200
	node.nextnode = node1 //将node和node1链接起来

	node2 := new(linklist)
	node2.data = 300
	node1.nextnode = node2 //将node1和node2链接

	node4 := new(linklist)
	node4.data = 400
	node3.nextnode = node4

	//读取链表
	nownode := node //从头结点开始

	for {
		if nownode != nil { //如果不为空值，则执行打印和寻找下一个节点
			fmt.Println(nownode.data)
			nownode = nownode.nextnode //将下一个节点赋给nownode
			continue

		}
		break

	}

}
