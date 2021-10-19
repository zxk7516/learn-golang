package main

import (
	"fmt"
	"reflect"
)

func main() {
	var y float64 = 3.4
	var x = &y

	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println("value: ", v, v.String())
	fmt.Println("value: ", t)

	*x = 3.2
	// v = reflect.ValueOf(x)
	// t = reflect.TypeOf(x)

	fmt.Println("value: ", v, v.String())
	fmt.Println("value: ", t)
}
