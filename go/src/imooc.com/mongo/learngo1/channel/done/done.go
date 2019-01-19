package main

import (
	"fmt"
	"sync"
)

const max int = 10

func doWorker(id int,
	worker worker) {
	for n := range worker.in {
		fmt.Printf("Worker %d received %d\n", id, n)
		worker.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, group *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			group.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(2 * max)

	var workers [max]worker
	for i := 0; i < max; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- i
	}

	for i, worker := range workers {
		worker.in <- i + max
	}

	wg.Wait()
}

func main() {
	chanDemo()

}
