package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	var a [3]int
	fmt.Printf("%v\n", a)
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("%v\n", a1)
	var b = make([]int, 0, 10)
	fmt.Printf("%v\n", b)
	var c = make([]int, 10)
	fmt.Printf("%v\n", c)

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

}

func main2313lkjkl() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main223kljld() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

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

func main211118() {
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
