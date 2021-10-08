package main

import (
	"fmt"
	"sync"
	"time"
)

// 并行，多个任务作用在多个cpu, **多任务多Cpu**
// 同一个时间点，有多个任务在同时运行
// 在微观角度，同一时间只有一个任务运行

// 并发, 多个任务在同一个cpu运行, **多任务单cpu**
// 在微观角度，同一时间只有一个任务运行

// Go 协程的特点
// 1. 有独立的栈空间
// 2. 共享程序堆空间
// 3. 调试由用户(主线程)控制
// 4. 协程是轻量级线程

// Go 协程调度模型 (M, P, G)
// M: 操作线索控制的主线程(物理线程)
// P: 协程执行需要的上下文
// G: 协程

// M0 正在执行G0协程，后面还有 G1, G2, G3
// G0 协和阻塞, M0将G1, G2, G3安排到线程程中的其它空闲线程 M1
// 如果没有 M1, 就创建线程
func test(tag string, max int) {
	for i := 0; i < max; i++ {
		fmt.Printf("%v hello,world, #%v\n", tag, i)
		time.Sleep(time.Microsecond * 700)
	}
}

func main() {

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		test("[co]", 10)
		wg.Done()
	}()
	// go test()   // test 和 main同时执行
	// test()     // test执行完成后，再执行main

	test("[main]", 5)
	wg.Wait()

}
