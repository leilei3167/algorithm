package main

import "fmt"

func bubblesort(array []int) {
	n := len(array)
	ifswap := true

	for i := n - 1; i > 0; i-- { //第一次比较n-1次，第二次n-2次以此类推
		for j := 0; j < i; j++ { //每一轮的比较次数
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				ifswap = false //发生交换就置为false
			}

		}
		if ifswap { //如果在第一轮都没有交换过 说明是有序数组
			return

		}

	}

}
func main() {
	a := []int{6, 4, 2, 1, 53, 2, 6}
	bubblesort(a)
	fmt.Println(a)
}
