package main

import (
	"fmt"
	"sync"
	"time"
)

func incrementCounter(c *Counter) {
	c.count++
}

func main() {

	counter := Counter{
		count: 0,
		mutex: &sync.Mutex{},
	}

	for i := 0; i < 1000; i++ {
		go incrementCounter(&counter)
	}

	time.Sleep(2 * time.Second)

	fmt.Println(counter.count)
}

type Counter struct {
	count int
	mutex *sync.Mutex
}
