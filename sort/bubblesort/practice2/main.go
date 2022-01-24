package main

import "fmt"

func bubble(arr []int) {
	tag := true
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tag = false
			}
			if tag { //第一次比较都没发生交换的话 说明是有序的
				return
			}
		}

	}

}

func main() {
	a := []int{65, 3, 24, 71, 23, 44}

	bubble(a)
	fmt.Println(a)
}
