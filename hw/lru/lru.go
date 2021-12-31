package lru

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Cache struct {
	max   int
	l     *list.List
	Call  func(key interface{}, value interface{})
	cache map[interface{}]*list.Element
	mu    *sync.Mutex
}

type Node struct {
	Key   interface{}
	Value interface{}
}

func New(maxEntries int) *Cache {
	return &Cache{
		max:   maxEntries,
		l:     list.New(),
		cache: make(map[interface{}]*list.Element),
		mu:    &sync.Mutex{},
	}
}

func (l *Cache) Add(key interface{}, val interface{}) error {
	if l.l == nil {
		return errors.New("lruCache not initied")
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if e, ok := l.cache[key]; ok {
		e.Value.(*Node).Value = val
		l.l.MoveToFront(e)
		return nil
	}
	ele := l.l.PushFront(&Node{
		Key:   key,
		Value: val,
	})

	l.cache[key] = ele
	if l.max != 0 && l.l.Len() > l.max {
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache, node.Key)
			if l.Call != nil {
				l.Call(node.Key, node.Value)
			}

		}
	}
	return nil
}

func (l *Cache) Get(key interface{}) (val interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if ele, ok := l.cache[key]; ok {
		l.l.MoveToFront(ele)
		return ele.Value.(*Node).Value, true
	}
	return
}

func LRU(operators [][]int, k int) []int {
	cache := New(k)
	res := make([]int, 0)
	for _, input := range operators {
		if input[0] == 1 {
			cache.Add(input[1], input[2])
		} else if input[0] == 2 {
			val, ok := cache.Get(input[1])
			if !ok {
				res = append(res, -1)
			} else {
				fmt.Println(val)
				v := val.(int)
				res = append(res, v)
			}
		}

	}
	return res
}
