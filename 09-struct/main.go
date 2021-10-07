package main

import (
	"fmt"
)

func main() {

	var cat Cat
	cat.Name = "MrBlack"
	cat.Age = 100
	cat.Color = "black"

	cat1 := Cat{Name: "MrWhite", Age: 3, Color: "white", Hobby: "Soccer"}
	fmt.Println(cat, cat1)

	fmt.Printf("%p, %p, %p, %p, %p\n", &cat1, &cat1.Name, &cat1.Age, &cat1.Color, &cat1.Hobby)

	p := Person{}
	fmt.Printf("%p, Name: %p, Age: %p, Scores: %p, ptr: %p, slice: %p, map1: %p, end: %p\n", &p, &p.Name, &p.Age, &p.Scores, &p.ptr, &p.slice, &p.map1, &p.end)

	var monster1 Monster
	monster1.Name = "abc"
	monster1.Age = 199

	monster2 := monster1
	fmt.Printf("%p, %p\n", &monster1, &monster2)

	monster3 := Monster{Name: "asdf", Age: 18}
	fmt.Println(monster3)
	monster4 := Monster{"WWiill", 18}
	fmt.Println(monster4)

	var monster5 *Monster = new(Monster)
	//var monster6 *Monster = &Monster{}  // same
	monster5.Name = "test"
	// (*monster5).Name = "test" because monster5 if a poinster of [Monster]
	// it can cat just like a [Monster]
	monster5.Age = 231
	fmt.Println(monster5)

	// 数组中的 Point 可以省略，因为已经定义了 [2]Point
	points := [2]Point{{1, 2}, {2, 4}} // [2]Point{Point{1,2}, Point{2,4}}
	rect := Rect{Point{1, 3}, Point{3, 4}}
	fmt.Println(points, rect)
	fmt.Printf("%p, %p, %p, %p\n\n", &rect.LeftUp.x, &rect.LeftUp.y, &rect.RightDown.x, &rect.RightDown.y)
	rect2 := Rect2{&Point{1, 3}, &Point{3, 4}}
	fmt.Printf("%p, %p, %p, %p\n\n", &rect2.LeftUp.x, &rect2.LeftUp.y, &rect2.RightDown.x, &rect2.RightDown.y)

	// 结构体转换，（字段名一样，个数一样，类型一样)
	px := Point{1, 3}
	vx := Vector(px)
	p2x := Point2(px)
	fmt.Println(vx)
	fmt.Println(p2x)
	fmt.Printf("%p ,%p, %p\n", &px, &vx, &p2x)
}

type Rect struct {
	LeftUp, RightDown Point
}

type Rect2 struct {
	LeftUp, RightDown *Point
}

type Point struct {
	x, y int
}

type Vector struct {
	x, y int
}

type Point2 Point

type Monster struct {
	Name string
	Age  int
}

// [pointer_to_string, uint8, pointer_to_string]
type Cat struct {
	Name  string // 16 memory 8byte poitner + 8 byte string lenght
	Age   uint64 // 8  uint8  same as uint64
	Color string // 16 memory 8byte poitner + 8 byte string lenght
	Hobby string // 16 memory 8byte poitner + 8 byte string lenght
}

// 结构体所有字段在内存中是连续的
type Person struct {
	Name   string            // 16
	Age    int               // 8
	Scores [5]float64        // 8 * 5
	ptr    *int              // 8
	slice  []int             // 8 * 3  pointer + cap + len
	map1   map[string]string // 8 pointer
	end    int
}
