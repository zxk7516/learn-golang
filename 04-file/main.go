package main

import (
	"bytes"
	rand2 "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// test2()
	// test_slice_on_array()
	generateRandomKey()
	return
	test_slice2()
	return
	file, err := os.Create("file.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("write file in golang\n")

	// Strings
	strings.EqualFold("abc", "Abc")
	strings.Index("NLT_001", "001") //  4
	strings.Index("NLT_001", "abc") // -1
	strings.Count("abcab", "abc")   // 1
	strings.Count("abcab", "ab")    // 2

}

var HelloError error = errors.New("hello")

func readConfig(i int) (s string, err error) {
	if i%2 == 0 {
		s = "even"
		// return "even",nil
	} else {
		// err = errors.New("hello")
		// return "", e

		err = HelloError
	}
	return
}

func monthDays(year, month int) (days int, err error) {
	if month < 1 || month > 12 {
		err = errors.New("wrong month")
		return
	}
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 12}
	days = months[month]
	if month == 2 && year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		days += 1
	}
	return
}

func printCal(year, month int) (days int, err error) {

	return

}

func test2() {
	var intArr [26]int
	arrLen := len(intArr)
	// 数组随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrLen; i++ {
		intArr[i] = rand.Intn(100)
	}

	// 反转数组
	var tmp int
	for i := 0; i < arrLen/2; i++ {
		tmp = intArr[i]
		intArr[i] = intArr[arrLen-i-1]
		intArr[arrLen-i-1] = tmp
	}
}

func test_slice_on_array() {
	// 切片是引用类型
	// 数组是值类型
	// 切片是指向数组的一个指针
	// 切片的长度变化时，有原数组空间不足的情况下会重新分配内存(堆内存)，并将切片指向新分配数组的地址
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:3] //       * pinter here [ 22,33 )
	fmt.Println(slice)   // [22,33]
	slice[0] = 222
	fmt.Println(intArr, slice) // [11,222,33,44,55] [222,33] // slice is still pointer to array
	fmt.Printf("->  %p %p\n", &intArr[1], slice)
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 4,  remainglint space of origin array

	slice = append(slice, 4)
	fmt.Printf("->4 %p %p\n", &intArr[1], slice)
	fmt.Println(intArr, slice)

	slice = append(slice, 5)
	fmt.Printf("->5 %p %p\n", &intArr[1], slice)
	slice = append(slice, 6)
	fmt.Printf("->6 %p %p\n", &intArr[1], slice)
	slice = append(slice, 7)
	slice[0] = 2
	fmt.Println(intArr, slice) // slice is not pointer to array,

}

func test_slice2() {
	// 切片必需在make后使用,只有栈上的元素可以用零值初始化
	slice := make([]int, 5)
	fmt.Println(slice, cap(slice), len(slice))
	fmt.Printf("%p", slice) // 0xc000016180

	slice = append(slice, 1)
	fmt.Println(slice, cap(slice), len(slice))
	fmt.Printf("%p", slice) // 0xc000106000 内存已经重新分配了
	for i, v := range slice {
		fmt.Println("i=", i, "v=", v)
	}
}

func getRandomPassword() {
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := generateRandomPassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	fmt.Println(password)
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generateRandomPassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder
	// set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}
	// set numeric
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}
	// set uppercase
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}
	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)

}

func generateRandomKey() {
	c := 10
	b := make([]byte, c)

	_, err := rand2.Read(b)
	if err != nil {
		fmt.Println("error generate failed")
		return
	}
	fmt.Println(b, bytes.Equal(b, make([]byte, c)))
}
