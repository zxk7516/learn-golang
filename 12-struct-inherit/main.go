package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("学生名=%v 年龄:%v 成绩:%v\n", stu.Name, stu.Age, stu.Score)

}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生考试中")
}

type Graduate struct {
	Student
	Major string
}

func (gra *Graduate) testing() {
	fmt.Println("大学生考试中")
}

type A struct {
	Name string
	Age  int
}

func (a *A) SayOk() {
	fmt.Println("A sayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello ", a.Name)
}

type B struct {
	a int
	A
	b int
}

type E struct {
	A
}

type F struct {
	int
	string
}
type C struct {
	*A // 匿名结构体
	*E // 匿名结构体
}
type C2 struct { // 多重继承, 尽量不要使用多重继承，也即，一个结构体中只有一个匿名结构体
	A // 匿名结构体
	E // 匿名结构体
}

type D struct {
	a A // 有名结构体
}

func main() {
	p := &Pupil{}
	p.Student.Name = "Tom"
	p.Student.Age = 8
	p.testing()
	p.Student.SetScore(70)
	p.Student.showInfo()

	g := &Graduate{}
	g.Student.Name = "Marry"
	g.Student.Age = 20

	g2 := &Graduate{
		Student: Student{Name: "Jim", Age: 15, Score: 89},
		Major:   "hello",
	}
	fmt.Println(g2)
	g.testing()
	g.Student.SetScore(60)
	g.Student.showInfo()

	var b B = B{}
	b.Name = "zxk"
	// B的结构僶中没有Name, 但是的嵌入的A结果体中有该字段，可以直接调用A中的字段
	// 如果B和A中有同名字段，则访问B中的字段,(就近访问)
	// 要特指A中的字段，需要有类型 B.A.{Field}
	b.Age = 28
	b.SayOk()
	b.hello()
	fmt.Printf("%p, %p, %p, %p, %p, %p", &b, &b.a, &b.A, &b.A.Name, &b.A.Age, &b.b)

	var c C
	c = C{&A{Name: "zxk", Age: 12}, &E{A: A{Age: 8, Name: "zy"}}}
	c = C{&A{Name: "zxk", Age: 12}, &E{A{Age: 8, Name: "zy"}}}
	c.Age = 9
	c.E.Age = 8
	c.A.Name = "tom"
	// c.Name // error, Name 就繼承自A還是B不清楚
	fmt.Println("c :", c)
	c2 := C2{A{Name: "zxk", Age: 12}, E{A{Age: 8, Name: "zy"}}}
	fmt.Println("c2 :", c2)
	c2.Age = 9
	c2.E.Age = 8
	c2.A.Name = "tom"
	fmt.Println("c2 :", c2)


	f := F{int:8,string:"abc"}
	fmt.Printf("%v = {int = %v, string = %v}\n", f, f.int, f.string)

}
