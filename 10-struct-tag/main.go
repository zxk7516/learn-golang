package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := Monster{"Zoe", 10, "Thunder"}
	fmt.Println(m)

	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json error", err)
	}
	fmt.Println(string(json))

	a := A{1}
	fmt.Printf("%v @ %p\n", a.Num, &a.Num)
	a.test()
	a.test2()

	b := B{2}
	A(b).test()

}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

type A struct {
	Num int
}
type B struct {
	Num int
}

// a 相当于一个可以改名的 this, 当没有*时，调用方法时，结构体会被复制到函数中
func (a A) test() {
	fmt.Printf("%v @ %p\n", a.Num, &a.Num)
}

// 有*时，调用方法时，结构体的指针或者引用会被引入到函数中
func (a *A) test2() {
	fmt.Printf("%v @ %p\n", a.Num, &a.Num)
}
