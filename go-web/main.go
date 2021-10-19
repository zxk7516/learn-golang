package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("workder:%d start job: %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("workder:%d end job: %d\n", id, j)
		results <- j * 2
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	wg.Add(runtime.NumCPU())
	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 48; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)
	for m := range results {
		fmt.Println("results: ", m)
	}

	// for {
	// 	m, ok := <-results
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("results: ", m)
	// }
}
