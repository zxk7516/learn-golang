package main

import (
	"fmt"
	"math/rand"
)

// 0 0 0 0 0 0
// 0 0 1 0 0 0
// 0 2 0 3 0 0
// 0 0 0 0 0 0

func main() {
	var arr [2][3]int
	arr[0][2] = 1
	arr[1][1] = 2

	fmt.Printf("arr[]      :%p\n", &arr)
	fmt.Printf("arr[0]     :%p\n", &arr[0])
	fmt.Printf("arr[1]     :%p\n", &arr[1])
	fmt.Printf("arr[0][0]  :%p\n", &arr[0][0])
	fmt.Printf("arr[0][1]  :%p\n", &arr[0][1])
	fmt.Printf("arr[1][0]  :%p\n", &arr[1][0])
	fmt.Printf("arr[1][1]  :%p\n", &arr[1][1])

	l1 := len(arr)
	l2 := len(arr[0])
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			fmt.Printf("[%d][%d] -> %d     ", i, j, arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println("遍历2")
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("[%d][%d] -> %d     ", i, j, v2)
		}
		fmt.Println()
	}
}

func test1() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

}
