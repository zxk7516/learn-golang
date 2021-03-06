package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int
}

func main() {
	// test2()
	rf1()
}

type Coder struct {
	Name string
}

func (c *Coder) String() string {
	return c.Name
}

func rf1() {
	coder := &Coder{Name: "Kippa"}
	val := reflect.ValueOf(coder)
	typ := reflect.TypeOf(coder)
	typeOfStringer := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(val)
	fmt.Println(typ)
	fmt.Println(typ.Kind())
	fmt.Println(val.Type().Implements(typeOfStringer))
}



// 
// 
//                             ┌───────────────┐ TypeOf/ValueOf──────────┐
//   ┌───────┐                 │               │ ◄─────────── │          │
//   │ value │◄────TypeCast────┤               │              │reflection│
//   └───────┘                 │interface value│  Interface() │  object  │
//                             │               │◄─────────────┤          │
//                             └───────────────┘              └──────────┘


func rf2() {

}

type user struct {
	UserId string
	Name   string
}

// 使用反射创建并操作结构体
func test2() {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)
	fmt.Println("reflect.TypeOf", st.Kind().String())
	st = st.Elem()
	fmt.Println("reflect.TypeOf.Elem", st.Kind().String())
	elem = reflect.New(st)

	fmt.Println("reflect.New", elem.Kind().String())
	fmt.Println("reflect.New.Elem", elem.Elem().Kind().String())

	model = elem.Interface().(*user)
	elem = elem.Elem()
	elem.FieldByName("UserId").SetString("1234")
	elem.FieldByName("Name").SetString("zxk")
	fmt.Println("model model.Name", model, model.Name)
}

type cal struct {
	Num1 int
	Num2 int
}

func (c *cal) getSub() int {
	return c.Num1 - c.Num2
}

func test03() {
	c := cal{1, 2}
	rf := reflect.ValueOf(c)

	fileC := rf.NumField()
	for i := 0; i < fileC; i++ {
	}

}
