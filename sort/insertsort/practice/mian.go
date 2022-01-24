package main

import "fmt"

func insertsort(m []int) {
	n := len(m)
	//n-1迭代
	for i := 1; i <= n-1; i++ {
		deal := m[i] //待排数，从第二个开始算
		j := i - 1   //待排左边的下标
		for ; j >= 0 && deal < m[j]; j-- {
			m[j+1] = m[j]

		}
		m[j+1] = deal

	}
	fmt.Println(m)
}
func main() {
	a := []int{3, 4, 1, 2}
	insertsort(a)
}
