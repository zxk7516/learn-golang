package main

import "fmt"

type BInterface interface{
test01()
}
type CInterface interface{
test02()
}

// 接口重载
type AInterface interface{
	BInterface
	CInterface
	test03()
}

type Stu struct{i int}

func (stu Stu) test01(){}
func (stu Stu) test02(){}
func (stu Stu) test03(){}

func main(){
	stu  := Stu{}
	var a AInterface
	a = stu
	var a2 AInterface
	a2 = &stu

	a.test01()
	a.test02()
	a.test03()

	a2.test01()
	a2.test02()
	a2.test03()

	stu.i = 5
	fmt.Printf("%p\n%p\n%p\n", &stu,&a,&a2)
	fmt.Printf("%P = %v\n%P = %v\n%P = %v\n", &stu,stu, &a, a,&a2,a2)

	// 空接口没有任何方法，可以接收任意数据类型
	receiveAnyType(&stu)
	receiveAnyType(stu)
}

func receiveAnyType(data interface{}) {
	fmt.Println(data)
}
