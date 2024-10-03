package main

import (
	"fmt"
	"sync"
	"time"
)

func printEven(cpe *sync.Cond, cpo *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		cpe.L.Lock()
		cpe.Wait()
		fmt.Println(i)
		cpo.Signal()
		cpe.L.Unlock()
	}
}

func printOdd(cpe *sync.Cond, cpo *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 99; i += 2 {
		cpo.L.Lock()
		cpo.Wait()
		fmt.Println(i)
		cpe.Signal()
		cpo.L.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	canPrintEven := sync.NewCond(&mutex)
	canPrintOdd := sync.NewCond(&mutex)

	wg.Add(2)

	go printEven(canPrintEven, canPrintOdd, &wg)
	go printOdd(canPrintEven, canPrintOdd, &wg)

	time.Sleep(time.Second)
	canPrintOdd.Signal()

	wg.Wait()
}
