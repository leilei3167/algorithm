package main

import (
	"fmt"
	"sync"
)

/* 队列先进先出，和栈操作顺序相反，我们这里只实现入队，和出队操作，其他操作和栈一样。
应着重考虑的是出队后的缩容问题 */
type arrayqueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

//入队
func (queue *arrayqueue) in(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//直接在切片中增加元素
	queue.array = append(queue.array, v)
	queue.size = queue.size + 1

}

//出队
func (queue *arrayqueue) out() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//判断是否已空
	if queue.size == 0 {
		return "队列已空"

	}
	//要取的就是第0个元素
	v := queue.array[0]
	//要考虑缩容
	/*    直接原位移动，但缩容后继的空间不会被释放
	      for i := 1; i < queue.size; i++ {
	          // 从第一位开始进行数据移动
	          queue.array[i-1] = queue.array[i]
	      }
	      // 原数组缩容
	      queue.array = queue.array[0 : queue.size-1]
	*/
	newarray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newarray[i-1] = queue.array[i]

	}
	queue.array = newarray
	return v
}
func (queue *arrayqueue) top() string {
	return queue.array[0]
}

func main() {
	a := new(arrayqueue)
	a.in("雷磊1")
	a.in("郝运2")
	a.in("杨真3")
	fmt.Println("下一个出来的是：", a.top())

}
