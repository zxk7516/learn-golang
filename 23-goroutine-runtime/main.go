package main

import (
	"fmt"
	"runtime"
)

func main() {

	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	maxProcs := runtime.GOMAXPROCS(20)
	fmt.Println(maxProcs)
}
