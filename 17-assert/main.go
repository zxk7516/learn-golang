package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}

	a = point
	var b Point

	//b = a //E: cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point) // 类型断言 => error
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y := x.(float32) // 类型断言时，如果类型不匹配，就会panic
	fmt.Printf("y: %T %v \n", y, y)

	// 带检测的类型断言
	y2, ok := x.(float64)
	// if y2,ok := x.(float64); ok {}
	if ok {
		fmt.Printf("转化成功 %v\n", y2)
	} else {
		fmt.Println("转化失败")
	}

	fmt.Println("end")
	var n1 float32 = 12.3
	var n2 float64 = 12.3
	var n3 int32 = 32
	var name string = "北京"
	TypeJudge(n1, n2, n3, name, b, &b)

}

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("#%v params type is bool ,value: %v\n", i+1, x)
		case float32:
			fmt.Printf("#%v params type is float32 ,value: %v\n", i+1, x)
		case float64:
			fmt.Printf("#%v params type is float64 ,value: %v\n", i+1, x)
		case int, int32, int64:
			fmt.Printf("#%v params type is int ,value: %v\n", i+1, x)
		case string:
			fmt.Printf("#%v params type is string ,value: %v\n", i+1, x)
		case Point:
			fmt.Printf("#%v params type is point ,value: %v\n", i+1, x)
		case *Point:
			fmt.Printf("#%v params type is *point ,value: %v\n", i+1, x)
		default:
			fmt.Printf("#%v params type is unknown, value: %v\n", i+1, x)
		}
	}
}
