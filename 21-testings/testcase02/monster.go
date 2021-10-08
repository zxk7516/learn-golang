package testcase02

import (
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() (err error) {
	// seralize to json
	// save to file

	data, err := json.Marshal(this)
	if err != nil {
		return
	}
	filepath := "monster.log"
	err = ioutil.WriteFile(filepath, data, 0666)
	return
}

func (this *Monster) Restore() (err error) {
	filepath := "monster.log"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, this)
	return
}
