package main

import (
	"fmt"
	"sort"
)

func main() {
	var m1 map[string]string
	m1 = make(map[string]string, 10)

	m1["no3"] = "Kate"
	m1["no1"] = "Tom"
	m1["no4"] = "Jef"
	m1["no2"] = "Jerry"
	m1["no6"] = "Jimmy"
	m1["no1"] = "Peter"
	m1["no5"] = "Walt"

	fmt.Println(m1)

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	heros := map[string]string{
		"no1": "Archer",
		"no2": "Hammer",
		"no3": "Rider",
	}
	fmt.Println(heros)

	students := make(map[string]map[string]string)
	students["no0001"] = map[string]string{"name": "张三", "sex": "male"}
	students["no0002"] = map[string]string{"name": "李四", "sex": "male"}

	students["no0003"] = make(map[string]string)
	students["no0003"]["name"] = "王芳"
	students["no0003"]["sex"] = "female"
	fmt.Println(students)
	for no, student := range students {
		fmt.Print(no, " {")
		for k, v := range student {
			fmt.Print(k, ":", v, ",")
		}
		fmt.Println("}")
	}
	delete(students, "no0001")
	delete(students, "no____")
	fmt.Println(students, "len:", len(students))
	// delete all keys, make new map
	students = make(map[string]map[string]string)

	students2 := make(map[string]Student)
	students2["no001"] = Student{name: "张三", sex: "male"}
	students2["no003"] = Student{name: "王芳", sex: "female"}
	students2["no002"] = Student{name: "李四", sex: "male"}
	fmt.Println(students2)

	for k, v := range students2 {
		fmt.Println(k, " ", v)
	}

	// map slice
	monsters := make([]map[string]string, 0)
	monsters = append(monsters, map[string]string{
		"name": "st",
		"age":  "500",
	})
	monsters = append(monsters, map[string]string{
		"name": "pa",
		"age":  "400",
	})
	fmt.Println(monsters)

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)
	modify(map1)

	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("map[%v] = %v\n", k, map1[k])
	}

	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["nickname"] = "xx"
	modifyUser(users, "tom")
	modifyUser(users, "smith")
	fmt.Println(users)
}

func modifyUser(map1 map[string]map[string]string, user string) {
	if map1[user] == nil {
		map1[user] = make(map[string]string, 2)
		map1[user]["pwd"] = "888"
		map1[user]["nickname"] = user
	} else {
		map1[user]["pwd"] = "888"
	}
}

// map 是引用类型
func modify(map1 map[int]int) {
	map1[10] = 900
}

type Student struct {
	name string
	sex  string
}
