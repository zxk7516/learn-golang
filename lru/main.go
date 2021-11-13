package main

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
	UseBytes() int
}

type entry struct {
	key string
}

func main() {

}
