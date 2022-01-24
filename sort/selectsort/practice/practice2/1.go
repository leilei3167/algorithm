package main

import "fmt"

func sort(m []int) {
	n := len(m)
	//进行n-1次迭代
	for i := 0; i < n-1; i++ {
		min := m[i]
		minindex := i
		for j := i + 1; j < n; j++ {
			if m[j] < min {
				min = m[j]
				minindex = j
			}
		}
		if i != minindex { //和选出的最小值交换
			m[i], m[minindex] = m[minindex], m[i]
		}

	}
	fmt.Println(m)
}
func main() {
	a := []int{5, 4, 3, 2, 1}
	sort(a)
}
