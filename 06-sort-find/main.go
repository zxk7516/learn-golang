package main

import "fmt"

func main() {
	arr := []int{24, 69, 80, 57, 13}
	bubbleSort(arr)
	fmt.Println(arr)
	fmt.Println(midFind(arr, 0))
	fmt.Println(midFind(arr, 57))
	fmt.Println(midFind(arr, -57))
	fmt.Println(midFind(arr, 80))

	fmt.Println(midFindRec(arr, 0))
	fmt.Println(midFindRec(arr, 57))
	fmt.Println(midFindRec(arr, -57))
	fmt.Println(midFindRec(arr, 80))
}

func midFindRec(arr []int, find int) (pos int) {
	return midFindRec_(arr, find, 0, len(arr)-1)
}
func midFindRec_(arr []int, find, start, end int) (pos int) {
	if start > end {
		pos = -1
		return
	}
	pos = (start + end) / 2
	if arr[pos] < find {
		return midFindRec_(arr, find, pos+1, end)
	} else if arr[pos] > find {
		return midFindRec_(arr, find, start, pos-1)
	} else {
		return
	}
}

func midFind(arr []int, find int) (pos int) {
	start := 0
	end := len(arr) - 1
	pos = (start + end) / 2
	for {
		if arr[pos] < find {
			start = pos + 1
		} else if arr[pos] > find {
			end = pos - 1
		} else {
			return
		}
		if start > end {
			break
		}
		pos = (start + end) / 2
	}
	pos = -1
	return
}

func bubbleSort(arr []int) {
	var tmp int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}
