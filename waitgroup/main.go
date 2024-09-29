package main

import (
	"fmt"
	"sync"
)

func printNumber(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()
}
