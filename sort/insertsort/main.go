package main

import "fmt"

func insertsort(m []int) {
	n := len(m)
	//进行n-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := m[i] //待排序的数
		j := i - 1   //待排序数左边的一个数

		//一直往左边找，比待排序数大的都往后挪，腾空位给待排序数插入
		for ; j >= 0 && deal < m[j]; j-- {
			m[j+1] = m[j] //后移

		}
		m[j+1] = deal

		fmt.Printf("第%d次比较结果%v\n", i, m)
	}

}
func main() {
	a := []int{3, 4, 1, 2}
	insertsort(a)
}
