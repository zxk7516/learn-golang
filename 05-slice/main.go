package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \t", i, slice[i])
	}
	fmt.Println()
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v \t", i, v)
	}
	fmt.Println()

	slice2 := arr[6:]
	slice_all := arr[:]
	slice3 := slice_all[:3]
	// 1.
	// 2. 切片直接定义是无法使用的, 需要先定义或者, 从已经初始化的数组或者切片上切出一个
	// 3. 切片还可以继续切
	fmt.Print(slice2, slice3)
	fmt.Println()
	fmt.Println(slice_all)

	str := "hello@email.com"
	fmt.Println(GetNameOfEmail(str))

	chineseName := "曾祥康"
	fmt.Println(chineseName)
	fmt.Println(GetFirstNameOfChineseName(chineseName))
	fmt.Println(HidenPartOfChineseName(chineseName))

	fmt.Println(fbn(20))

}

func GetNameOfEmail(str string) (name string) {
	atIndex := strings.Index(str, "@")
	strNameSlice := str[:atIndex]
	name = string(strNameSlice)
	return
}

func GetFirstNameOfChineseName(name string) string {
	slice := []rune(name)
	return string(slice[0])
}

func HidenPartOfChineseName(name string) string {
	slice := []rune(name)
	slice[1] = '某'
	return string(slice)
}

func fbn(n uint64) []uint64 {
	slice := make([]uint64, n)
	slice[0] = 1
	slice[1] = 1
	var i uint64
	for i = 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
