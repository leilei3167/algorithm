package main

import "fmt"

func bubble(arr []int) {
	/* 冒泡排序，第一次进行n-1次比较，第二次进行n-2次比较，以此类推 */
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}
	fmt.Println(arr)

}
func main() {
	a := []int{45, 21, 41, 1, 23, 77, 2, 3, 4}
	bubble(a)
}
