package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type Foo struct {
	A int8 // 1
	B int8 // 1
	C int8 // 1
}

type Bar struct {
	x int32 //         4
	// padding 4
	y *Foo // 8
	z bool //         1
	// padding 7
}
type Bar2 struct {
	x int32 //         4
	z bool  //         1
	// padding 3
	y *Foo //         8
}

func main() {
	var f Foo
	fmt.Println(unsafe.Sizeof(f))
	fmt.Printf("A: %p, B: %p, C: %p\n", &f.A, &f.B, &f.C)

	var b1 Bar
	fmt.Println(unsafe.Sizeof(b1))
	fmt.Printf("x: %p, y: %p, z: %p\n", &b1.x, &b1.y, &b1.z)

	// m()

	m2()

}

type job struct {
	value int64
}

type result struct {
	job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg = &sync.WaitGroup{}

func zhoulin(zl chan<- *job) {
	defer wg.Done()
	for i := 0; i < 48; i++ {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
	close(jobChan)
}

func baodelu(zl <-chan *job, res chan<- *result) {
	defer wg.Done()
	for {
		j, ok := <-zl
		if !ok {
			break
		}
		sum := int64(0)
		n := j.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    *j,
			result: sum,
		}
		res <- newResult
	}

}

func m() {
	wg.Add(1)
	go zhoulin(jobChan)

	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}
	go func() {
		for result := range resultChan {
			fmt.Printf("value:%d sum:%d\n", result.job.value, result.result)
		}
	}()
	wg.Wait()
}

func m2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func server() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d: %d\n", i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	var q [3]int = [3]int{1, 2, 3}
	q = [...]int{1, 2, 3} // eq
	fmt.Println(q)

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
	JPY
)

func StringStartWith(s, prefix string) bool {
	return len(s) >= len(prefix) &&
		s[:len(prefix)] == prefix
}

func StringEndWith(str, suffix string) bool {
	len_str := len(str)
	len_suf := len(suffix)

	return len_str >= len_suf &&
		str[len_str-len_suf:] == suffix
}

var StringContainsO = strings.Contains

func StringContains(str, substr string) bool {
	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], substr) {
			return true
		}
	}
	return false
}

func tt2() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "C￥", JPY: "J￥"}
	fmt.Println(symbol[USD])
}
