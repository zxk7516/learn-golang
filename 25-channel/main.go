package main

import (
	"fmt"
	"sync"
)

func main() {

	numChan := make(chan int, 1999)
	resChan := make(chan int, 1999)
	exitChan := make(chan bool)
	go func(numChan chan<- int) {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}(numChan)

	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(numChan <-chan int, resChan chan<- int, exitChan chan<- bool) {
			for {

				n, ok := <-numChan
				if !ok {
					break
				}
				res := 0
				for m := 1; m <= n; m++ {
					res += m
				}
				resChan <- res
			}
			// exitChan <- true
			wg.Done()
		}(numChan, resChan, exitChan)
	}
	wg.Wait()
	close(resChan)
	close(exitChan)
	// wg.Wait()
	// close(resChan)
	for {
		n, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println("res: ", n)
	}

}

/*
func WriteData_test_chan4(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}
func readData_test_chan4(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read data : ", v)
	}
	exitChan <- true
}

func test_chan4() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go WriteData_test_chan4(intChan)
	go readData_test_chan4(intChan, exitChan)
	<-exitChan

}

func test_chan3() {
	inChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 200; i++ {
			inChan <- i
		}
		close(inChan)

	}()
	go func() {
		for i := range inChan {
			fmt.Println("i=", i)
		}
		wg.Done()
	}()
	wg.Wait()

}

func test_chann2() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2

	close(intChan)
	// intChan <- 3 // error 关闭的管道不能写，可以读
	i := <-intChan
	i2 := <-intChan
	fmt.Println(i, i2)

	// channel 支持 for--range 遍历
	// 1. 遍历时如果 channel 没有关闭，会出现  deadlock 错误
	// 2. 遍历时如果 channel 关闭了，会正常遍历，数据取完后退出循环
	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2) // 如果没有这一行，取到198时会报 deadlock 错误
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	fmt.Println("ok")

}

func test_chann() {

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "Flash"
	m2["hero2"] = "Hamer"

	mapChan <- m1
	mapChan <- m2

	<-mapChan
	<-mapChan

	catChan := make(chan interface{}, 4)
	catChan <- Cat{"小花猫", 4}

	cat := <-catChan
	a := cat.(Cat)
	fmt.Println(a)

}

type Cat struct {
	Name string
	Age  int
}
*/
