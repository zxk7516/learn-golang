package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println("小猴子在爬树")
}
func (this *LittleMonkey) flying() {
	fmt.Println("小猴子在飞翔")
}
func (this *LittleMonkey) swimming() {
	fmt.Println("小猴子在游泳")
}

type IClimb interface {
	climbing()
}
type ISwim interface {
	swimming()
}
type IFly interface {
	flying()
}

type LittleMonkey struct {
	Monkey
}

func main() {
	intSlice := []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)

	arr := [5]int{0, -1, 10, 7, 90}
	sort.Ints(arr[:])

	fmt.Println(intSlice, arr)

	monkey := &LittleMonkey{
		Monkey{Name: "Tom"},
	}

	monkey.climbing()

	var ifly IFly = monkey
	ifly.flying()

	var iswim ISwim = monkey
	iswim.swimming()

}
