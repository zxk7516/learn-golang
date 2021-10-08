package testcase02

import (
	"testing"
)

// 测试Store
func TestStore(t *testing.T) {
	var err error
	monster := Monster{Name: "红孩子",Age: 10,Skill: "喷火"}
	err = monster.Store()
	if err != nil {
		t.Fatal(err)
	}

}
// 测试Store
func TestRestore(t *testing.T) {
	var err error
	var monster Monster
	err = monster.Restore()
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(monster)

}
