package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	lock  *sync.Mutex
	l     *list.List
	cache map[interface{}]*list.Element
	max   int
}

type Node struct {
	Key   interface{}
	Value interface{}
}

func New(max int) *LRUCache {
	return &LRUCache{
		max:   max,
		l:     list.New(),
		cache: make(map[interface{}]*list.Element),
		lock:  new(sync.Mutex),
	}
}

func (l *LRUCache) Add(key interface{}, value interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if e, ok := l.cache[key]; ok {
		e.Value.(*Node).Value = value
		l.l.MoveToFront(e)
		return nil
	}
	ele := l.l.PushFront(&Node{
		Key:   key,
		Value: value,
	})
	l.cache[key] = ele
	if l.max != 0 && l.l.Len() > l.max {
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache, node.Key)
		}
	}
	return nil
}

func (l *LRUCache) Get(key interface{}) (val interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	l.lock.Lock()
	defer l.lock.Unlock()

	if ele, ok := l.cache[key]; ok {
		l.l.MoveToFront(ele)
		return ele.Value.(*Node).Value, true
	}
	return
}
