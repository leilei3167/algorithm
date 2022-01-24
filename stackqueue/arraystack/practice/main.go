package main

import "fmt"

/* 用数组实现栈 */
type arraytack struct {
	array []interface{}
	size  int
}

//入栈
func (stack *arraytack) push(v interface{}) {
	stack.array = append(stack.array, v)
	stack.size = stack.size + 1

}

//出栈
func (stack *arraytack) out() interface{} {
	if stack.size == 0 {
		return "没有可用元素！！！"

	} else {
		//要取出的元素下标就是size-1
		v := stack.array[stack.size-1]
		//缩容数组
		newarray := make([]interface{}, stack.size-1, stack.size-1)
		for i := 0; i < stack.size-1; i++ {
			newarray[i] = stack.array[i]

		}
		stack.array = newarray
		stack.size = stack.size - 1
		return v
	}

}

func (stack *arraytack) list() interface{} {
	if stack.size == 0 {
		return "没有可用元素！！！"

	} else {

		v := stack.array[stack.size-1]

		return v
	}

}

func main() {
	a := new(arraytack)
	a.push(45)
	a.push("leilei")
	a.push("55.5555")
	a.out()
	fmt.Println(a.list())
}
