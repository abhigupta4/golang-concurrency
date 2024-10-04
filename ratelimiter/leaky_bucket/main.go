package main

import (
	"context"
	"sync"
	"time"
)

func refreshToken(bucket chan int, tokens int) {
	for {
		time.Sleep(time.Second)
		for i := 0; i < tokens; i++ {
			bucket <- 1
		}
	}
}

func doWork(id int, bucket chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		select {
		case <-ctx.Done():
			println("Worker", id, "rate limited")
		case <-bucket:
			println("Worker", id, "work done")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	allowed := 3
	bucket := make(chan int, 100)
	go refreshToken(bucket, allowed)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(i, bucket, wg)
	}

	wg.Wait()
}
