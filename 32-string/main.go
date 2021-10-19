package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 8, 1, 4}
	res := quickSort(arr)
	fmt.Println(res)
	quickSortOnLocal(arr)
	fmt.Println(arr)

}

func bfs() {

}

func quickSort(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func quickSortOnLocal(arr []int) {
	l := len(arr)
	if l < 2 {
		return
	}
	pivot := quickSortOnLocalPivot(arr, 0, l-1)
	quickSort(arr[:pivot])
	quickSort(arr[pivot+1:])
}

func quickSortOnLocalPivot(arr []int, lo, hi int) int {
	x := arr[lo]
	j := hi + 1
	for i := hi; i > lo; i-- {
		if arr[i] >= x {
			j--
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[lo], arr[j-1] = arr[j-1], arr[lo]
	return j - 1
}
