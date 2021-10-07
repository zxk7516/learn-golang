package model

import "fmt"

type person struct {
	name string
	age  int
	sal  float64
}

func New_person(name string) *person {
	return &person{
		name: name,
	}
}

func (per *person) SetAge(age int) {
	if age < 0 || age > 150 {
		fmt.Println("age invalid")
	}
	per.age = age
}

func (per *person) GetAge() int {
	return per.age
}

func (per *person) SetSal(sal float64) {
	if sal < 3000 || sal > 30000 {
		fmt.Println("salary invalid")
	}
	per.sal = sal
}

func (per *person) GetSal() float64 {
	return per.sal
}
func (per *person) GetName() string {
	return per.name
}
