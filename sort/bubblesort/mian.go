package main

import "fmt"

/* 冒泡排序 */
//创建一个方法，对于传入的切片用冒泡排序
func Bubblesor(list []int) { //因为切片是引用类型，传入的是地址，这里排序的操作直接作用于原切片指向的数组
	n := len(list)

	didSwap := false

	//进行n-1次比较
	for i := n - 1; i > 0; i-- {
		//每次从第一位开始比较，比较到第i位就不比较了，因为前一轮最后一位已确定
		//第一次比较次数为n-1，第二次比较次数为n-2（确定了一位的顺序）
		for j := 0; j < i; j++ { //从切片第一个和第二个开始比较，如果前一个大则交换顺序
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}

		}
		//如果在一轮中没有交换过，那么说明是一个有序数组，已经排好序了，直接返回
		if !didSwap {
			return

		}
	}

}

func main() {
	a := []int{56, 12, 34, 1, 35, 67, 89, 98, 42}
	Bubblesor(a)
	fmt.Println(a)
}
