package main

// 最优的单向冒泡排序
func bubbleSort(a []int) {
	for n := len(a); n > 1; {
		newn := 0 // 用来保证一定退出最外层循环
		for i := 0; i < n; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				newn = i // 用来跳过比较已排序的元素
			}
		}
		n = newn
	}
}
