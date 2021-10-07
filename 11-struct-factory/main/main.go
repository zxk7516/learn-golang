package main

import (
	"fmt"
	"zxk/11-struct-factory/model"
)

func main() {
	var Stu = model.Student{
		Name:  "Tom",
		Score: 78.5,
	}
	fmt.Println(Stu)

	var stu = model.New_student("Tom~", 88.8)
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.Get_score())

	p := model.New_person("Zxk")
	p.SetAge(29)
	p.SetSal(6000.88)
	fmt.Println(p)
	fmt.Println(p.GetName(), "age:", p.GetAge(), "sal:", p.GetSal())

}
