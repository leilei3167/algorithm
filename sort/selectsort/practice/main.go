package main

import "fmt"

func selectsort(m []int) {
	n := len(m)

	for i := 0; i < n-1; i++ {
		min := m[i]
		minindex := i
		for j := i + 1; j < n; j++ { //选出最小值
			if m[j] < min { //如果有比第前几个还小的值，则将它赋给min
				min = m[j]
				minindex = j

			}

		}
		//交换
		if i != minindex {
			m[i], m[minindex] = m[minindex], m[i]
		}

	}
	fmt.Println(m)

}
func main() {
	a := []int{4, 3, 2, 1}
	selectsort(a)
}
