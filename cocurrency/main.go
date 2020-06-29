package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	result := CocurrencyWG()
	fmt.Println(result)
	fmt.Println("====================================================")
	result = CocurrencyChan()
	fmt.Println(result)
}

func CocurrencyWG() []int {
	work := make([]int, 100)
	result := make([]int, 100)
	for i := 0; i < len(work); i++ {
		work[i] = i
	}
	// split workers
	wg := &sync.WaitGroup{}
	numWorkers := 3
	size := (len(work) / numWorkers) + 1
	for i := 0; i < numWorkers; i++ {
		max := min((i+1)*size, len(work)-1)
		go doWorkWG(i, i*size, max, work, result, wg)
	}
	// wait for work to be done
	wg.Wait()
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func doWorkWG(id int, startI, endI int, work []int, out []int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for i := startI; i <= endI; i++ {
		out[i] = work[i] * work[i]
	}
}

func CocurrencyChan() []int {
	work := make(chan int, 20)
	result := make(chan int, 100)
	resultArray := []int{}
	// create work
	go func() {
		for i := 0; i < 100; i++ {
			work <- i
		}
		close(work)
	}()
	//

	for i := 0; i < 3; i++ {
		go doWorkChan(i, work, result)
	}
	for r := range result {
		resultArray = append(resultArray, r)
	}
	return resultArray
}

func doWorkChan(id int, work <-chan int, out chan<- int) {
	var w int
	for done := false; !done; w, done = <-work {
		out <- w * w
	}
	time.Sleep(1 * time.Second)
	// close(out)
}

// Go runs the given function in a goroutine and catches + logs panics. More
// advanced use cases should copy this implementation and modify it.
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				fmt.Printf("goroutine panic: %v\n%s", err, stack)
			}
		}()
		f()
	}()
}
