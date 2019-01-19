package main

import (
	"fmt"
	"time"
)

const max int = 10

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(-1, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(-1, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	//close(c)
	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [max]chan int
	for i := 0; i < max; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < max; i++ {
		channels[i] <- i
	}

	for i := 0; i < max; i++ {
		channels[i] <- i + max
	}
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
