package main

import "fmt"

/* 实现哈希表的增删改查 */

/* 思路：1.哈希表本质为数组链表，哈希表包含一个数组，数组的值指向一个链表的起点 */
type Hashtable struct {
	array [16]*Node
	cap   int //总共装了多少键值对

}
type Node struct {
	key   string
	value interface{}
	next  *Node
}

//初始化一个哈希表，数组16个值默认都是nil，cap默认是0
func Newhashtable() *Hashtable {
	return new(Hashtable)

}

/* 2.选定一个哈希函数，以便于计算存放的数组的标号 */
func hash(key string) int {
	index := 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}

//显示当前键值对的数量
func (this *Hashtable) Count() {
	fmt.Printf("当前已存储的键值对数量为%d\n", this.cap)

}

//加入键值对
func (this *Hashtable) Put(k string, v interface{}) {
	//首先先计算要放入的数组标号
	index := hash(k)
	//判断该标号是否是nil值，是就说明还没有放入过元素
	if this.array[index] == nil {
		//创建新的节点，并让标号指向他
		n := new(Node)
		n.key = k
		n.value = v
		this.array[index] = n
		fmt.Println("已完成数值添加！", k, v)

	} else { //就说明该下标已经有了元素，遍历这个链表看是否有k相同的 有的话就应该更新值，没有的话接入尾部
		firstnode := this.array[index] //取出链表头
		for firstnode != nil {
			if firstnode.key == k { //说明是在已有key上更新
				firstnode.value = v
				fmt.Printf("已将'%s'的值更新为：%v！\n", k, v)
				return
			}
			firstnode = firstnode.next

		} //跳出循环表示到末尾
		n := new(Node)
		firstnode = n
		n.value = v
		n.key = k
		fmt.Println("已完成数值添加:", k, ":", v)

	}
	this.cap++

}

//根据key删除值
func (this *Hashtable) Del(k string) {
	//根据key查找
	//先查找
	//先判断是否为空表
	if this.cap == 0 {
		panic("目前没有任何数据！请先添加数据！")
	}
	index := hash(k)
	//遍历该标号数组的链表

	for this.array[index] != nil {
		if this.array[index].key == k {
			fmt.Printf("已完成删除'%s'：%v\n", this.array[index].key, this.array[index].value)
			this.array[index] = this.array[index].next
			this.cap--
			return
		}
		this.array[index] = this.array[index].next
	}
	//出循环就说明没有找到对应的key
	fmt.Printf("没有找到\"%s\"对应的值！！\n", k)
	return

}

//根据key查找
func (this *Hashtable) Look(k string) interface{} {
	//先判断是否为空表
	if this.cap == 0 {
		panic("目前没有任何数据！请先添加数据！")
	}
	index := hash(k)
	//遍历该标号数组的链表

	for this.array[index] != nil { //此处如果用.next的话，则指向尾部之后就不会执行循环体
		if this.array[index].key == k {
			fmt.Printf("%s对应的值为：%v\n", this.array[index].key, this.array[index].value)
			return this.array[index].value
		}
		this.array[index] = this.array[index].next
	}
	//出循环就说明没有找到对应的key
	fmt.Printf("没有找到\"%s\"对应的值！！\n", k)
	return "没有该key对应的值！！"

}

//遍历哈希表
func (this *Hashtable) Read() {
	for i := 0; i < 16; i++ { //遍历哈希表，对于不是空值的 进一步遍历链表
		if this.array[i] != nil {

			for this.array[i] != nil { //遍历链表
				fmt.Printf("key：%s 对应的值是：%v\n", this.array[i].key, this.array[i].value)
				this.array[i] = this.array[i].next

			}

		}
	}
}

//test
func main() {
	a := Newhashtable()
	a.Put("1111", "盖伦")
	a.Put("2", 121)
	a.Put("daw", 1.11)
	a.Put("78", 1132)
	a.Put("111", "英雄联盟")
	a.Put("543", 3.14159)
	a.Put("1", "哈哈哈哈")
	a.Put("12", "丢雷") //增
	a.Del("12")       //删
	a.Look("12")      //查
	a.Put("12", "你好") //增
	a.Put("12", "我是你爹")
	a.Look("12")
	a.Del("12") //删
	a.Look("12")
	a.Del("1111")
	a.Del("2")
	a.Del("daw")
	a.Del("78")
	a.Del("543")
	a.Put("wo", 11)
	fmt.Println("-----------遍历一下----------")
	a.Count()
	a.Read()
}
