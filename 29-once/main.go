package main

import (
	"image"
	"sync"
	"fmt"
)

var icons map[string]image.Image

func loadIcon(path string) image.Image {
	return nil
}

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"right": loadIcon("right.png"),
		"up":    loadIcon("up.png"),
		"down":  loadIcon("down.png"),
	}
}

var wg = &sync.WaitGroup{}
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	once.Do(func() {
		close(ch1)
	})
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)

	wg.Add(3)
	go f1(a)
	go f2(a, b)

	var m = sync.Map{}
	m.Store("abc", 1)
	m.Load("abc")
	v, _ := m.LoadOrStore("abc", 3)
	fmt.Println(v)

}
