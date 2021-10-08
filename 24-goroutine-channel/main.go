package main

import (
	"fmt"
	"sync"
	_ "time"
)

// 计算1-200 的各个数的阶乘,放入到map中

// 思路
// 1. 编写函数，计算各个数的阶乘，并放入 map中
// 2. map 是全局的

var (
	mymap   = make(map[int]int, 200)
	wg      = &sync.WaitGroup{}
	ch      = make(chan [2]int, 20)
	mapLock = &sync.Mutex{}
)

// go build -race main.go
func test(n int) {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	ch <- [2]int{n, res}
	wg.Done()

	// mapLock.Lock()
	// defer mapLock.Unlock()
	// mymap[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		wg.Add(2)
		go test(i)
	}
	go func() {
		for {
			n := <-ch
			mymap[n[0]] = n[1]
			wg.Done()
		}
	}()

	wg.Wait()

	// time.Sleep(10*time.Second)
	// mapLock.Lock()
	// defer mapLock.Unlock()

	for i, v := range mymap {
		fmt.Printf("map[%3d] = %20d\n", i, v)
	}

}
