package main

import (
	"sync"
	"time"
)

func main() {
	semaphore := make(chan int, 3)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(i, semaphore, &wg)
	}

	wg.Wait()
}

func worker(id int, semaphore chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	semaphore <- 1

	time.Sleep(time.Second)
	println("Worker ", id)
	<-semaphore
}
