package main

import (
	"fmt"
	"sync"
)

/* 主要通过可变长数组来实现 */
type Arraystack struct {
	array []string
	size  int
	lock  sync.Mutex
}

//入栈
func (stack *Arraystack) push(v string) {
	stack.lock.Lock()         //将元素入栈，会先加锁实现并发安全。
	defer stack.lock.Unlock() //上锁保证并发安全

	stack.array = append(stack.array, v) //放入切片中
	stack.size = stack.size + 1

}

//出栈
func (stack *Arraystack) pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//先判断栈中是否还有数据
	if stack.size == 0 {
		panic("没有可供取出的元素！")
	}
	// 栈顶元素赋值给v
	v := stack.array[stack.size-1]
	//要考虑将栈空间缩小，即把切片缩容
	/* 方法一：stack.array=stack.array[0:stack.size-1] */
	//方法二：创建新数组
	newarray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newarray[i] = stack.array[i]
	}
	stack.array = newarray
	//栈中元素减1
	stack.size = stack.size - 1
	return v

}

//获取栈顶元素但不出栈
func (stack *Arraystack) peek() string {
	//先判断是否已空
	if stack.size == 0 {
		panic("没有元素！")

	}
	v := stack.array[stack.size-1]

	return v

}

//获取当前栈的大小
func (stack *Arraystack) storage() int {
	v := stack.size
	return v

}

//判断当前是否为空栈，为空的话返回true
func (stack *Arraystack) is() bool {
	if stack.size == 0 {
		return true

	} else {
		return false
	}

}

func main() {
	stu := new(Arraystack)
	stu.push("雷磊1")
	stu.push("郝运2")
	stu.push("杨真3") //入栈
	fmt.Println("此时栈尺寸", stu.storage())
	fmt.Println("此时栈顶元素为", stu.peek())
	//出栈
	fmt.Println("第一个出栈的是：", stu.pop())
	fmt.Println("第二个出栈的是：", stu.pop())
	fmt.Println("第三个出栈的是：", stu.pop())

}
