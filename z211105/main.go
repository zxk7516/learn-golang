package main

import "fmt"
import "sync"

type Db struct {
	name string
}

var once sync.Once
var db *Db

func initDb() *Db {
	return &Db{name: "test"}
}

func main() {
	a := foo(1, 2)
	fmt.Println(a)
}

func foo(a, b int) int {
	once.Do(func() {
		fmt.Println("hello")
	})
	return a + b
}
func bar(a, b int) int {
	return a - b
}
