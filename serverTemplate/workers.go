package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

type Worker struct {
	ID     int
	input  <-chan string
	output chan<- string
	done   chan bool
	once   sync.Once
	wg     *sync.WaitGroup
}

// Start is what should be called as the goroutine function.
// It ends the goroutine when the done channel is closed.
func (w *Worker) Start() {
	w.wg.Add(1)
	defer w.wg.Done()
	fmt.Printf("Starting workers %d\n", w.ID)
	// recover if panic inside goroutine
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("PANICKING in Worker %d: %+v\n%s", w.ID, r, debug.Stack())
			fmt.Println(err)
		}
	}()
	// loop until done is closed or input is closed
	for {
		select {
		case s, ok := <-w.input:
			if !ok {
				fmt.Printf("Closing worker %d because of closed input\n", w.ID)
				return
			}
			// actual work
			w.output <- fmt.Sprintf("Worker %d: processing %s", w.ID, s)
		case <-w.done:
			fmt.Printf("Closing worker %d because of closed done\n", w.ID)
			return
		}
	}
}

// Stop stops the worker by closing the done channel for the workers.
func (w *Worker) Stop() {
	w.once.Do(func() {
		close(w.done)
	})
}

func createWorkers(num int, input <-chan string, output chan<- string, wg *sync.WaitGroup) []*Worker {
	workers := make([]*Worker, num)
	for i := 0; i < num; i++ {
		workers[i] = &Worker{
			ID:     i,
			input:  input,
			output: output,
			done:   make(chan bool, 1),
			wg:     wg,
		}
	}
	return workers
}

func workers() {
	wg := &sync.WaitGroup{}
	input := make(chan string, 10)
	output := make(chan string, 10)
	workers := createWorkers(runtime.NumCPU()/2, input, output, wg)
	for i := range workers {
		go workers[i].Start()
	}
	workers[0].Stop()
	go func() {
		for _, s := range strs {
			input <- s
		}
		// close(input)
		for i := range workers {
			workers[i].Stop()
		}
		wg.Wait()
		close(output)
	}()
	for open, out := true, ""; open; out, open = <-output {
		if !open {
			break
		}
		fmt.Println(out)
	}
	wg.Wait()
}

var strs = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"11",
	"12",
	"13",
	"14",
	"15",
	"16",
	"17",
	"18",
	"19",
	"20",
	"21",
	"22",
	"23",
	"24",
	"25",
	"26",
	"27",
	"28",
	"29",
	"30",
	"31",
	"32",
	"33",
	"34",
	"35",
	"36",
	"37",
	"38",
	"39",
	"40",
	"41",
	"42",
	"43",
	"44",
	"45",
	"46",
	"47",
	"48",
	"49",
	"50",
	"51",
	"52",
	"53",
	"54",
	"55",
	"56",
	"57",
	"58",
	"59",
	"60",
	"61",
	"62",
	"63",
	"64",
	"65",
	"66",
	"67",
	"68",
	"69",
	"70",
	"71",
	"72",
	"73",
	"74",
	"75",
	"76",
	"77",
	"78",
	"79",
	"80",
	"81",
	"82",
	"83",
	"84",
	"85",
	"86",
	"87",
	"88",
	"89",
	"90",
	"91",
	"92",
	"93",
	"94",
	"95",
}
