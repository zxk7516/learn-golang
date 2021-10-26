package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := "abc"
	ss(s1)

}

func ss(s string) {
	fmt.Printf("%d", unsafe.Sizeof(s))

}
