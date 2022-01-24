package main

import "fmt"

/* 数组创建队列 */
type arrayqueue struct {
	array []interface{}
	size  int
}

//入队
func (queue *arrayqueue) Push(v interface{}) {
	queue.array = append(queue.array, v)
	queue.size++

}

//出队
func (queue *arrayqueue) Pop() interface{} {
	if queue.size == 0 {
		return "没有可取的元素！"

	}
	//要出的元素就是arr[0]
	v := queue.array[0]
	//缩容
	newarray := make([]interface{}, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newarray[i-1] = queue.array[i]
	}
	queue.array = newarray
	queue.size--
	return v

}
func (queue *arrayqueue) read() {
	//遍历
	if queue.size == 0 {
		panic("没有元素！")
	}
	for k, v := range queue.array {
		fmt.Printf("queue[%d]=[%v]\n", k, v)

	}
}

func main() {
	a := new(arrayqueue)
	a.Push(23)
	a.Push(23.22)
	a.Push("you")
	a.Push("love ")
	a.Push("me")
	a.read()
	fmt.Println("--------------出队一个后！-----------------")
	a.Pop()
	a.read()
}
