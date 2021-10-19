package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func main() {

	fmt.Println("================serilize================")
	monster := Monster{
		Name:     "牛魔王",
		Age:      15,
		Birthday: "1992-10-04",
		Sal:      1234.4,
		Skill:    "Sword",
	}

	data, _ := json.Marshal(monster)
	fmt.Println(string(data))

	a := make(map[string]interface{})
	a["Name"] = "牛魔王"
	a["Age"] = 15
	a["Birthday"] = "1992-10-04"
	a["Sal"] = 1234.4
	a["Skill"] = "Sword"
	data, _ = json.Marshal(a)
	fmt.Println(string(data))

	var slice = []Monster{
		Monster{"abc", 15, "1992-10-4", 1.1, "Sword"},
		{"def", 23, "2012-10-4", 2.1, "Stone"},
		{"hksdlaf", 9923, "2123-10-4", 98.3, "Gun"},
	}
	data, _ = json.Marshal(slice)
	fmt.Println(string(data))

	slice2 := [](map[string]interface{}){
		map[string]interface{}{"abc": 1, "def": a},
		map[string]interface{}{"qdaf": 1, "alskjdj": slice},
	}
	data, _ = json.Marshal(slice2)
	fmt.Println(string(data))
	fmt.Println("================deserilize================")

	jsonStr := "{\"name\":\"牛魔王\",\"age\":15,\"birthday\":\"1992-10-04\",\"sal\":1234.4,\"skill\":\"Sword\"}"
	var mon Monster
	json.Unmarshal([]byte(jsonStr), &mon)
	fmt.Println(mon)

	var m map[string]interface{} // 反序列化map时不需要make, Unmarshal函数会初始化
	// m = make(map[string]interface{})
	rv := reflect.ValueOf(&m)
	v := rv.Interface()
	m2 := v.(*map[string]interface{})
	fmt.Printf("%p %p, %T, %T,%T\n", &m, m2, m, m2, rv)
	// m3 := make(map[string]interface{},5)
	// m3["abc"] = 1
	// m2["abc"] = 1

	k := rv.Kind()
	fmt.Println(k, k.String(), rv.IsNil())
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m)

	jsonStr = "[{\"name\":\"abc\",\"age\":15,\"birthday\":\"1992-10-4\",\"sal\":1.1,\"skill\":\"Sword\"},{\"name\":\"def\",\"age\":23,\"birthday\":\"2012-10-4\",\"sal\":2.1,\"skill\":\"Stone\"},{\"name\":\"hksdlaf\",\"age\":9923,\"birthday\":\"2123-10-4\",\"sal\":98.3,\"skill\":\"Gun\"}]"

	var ss []Monster //  不需要初始化，原理同上
	json.Unmarshal([]byte(jsonStr), &ss)
	fmt.Println(ss)


}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return "??"
	}

	return ""
}
