package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go runScheduledTask(3*time.Second, printHi, &wg)

	wg.Add(1)
	go runScheduledTask(time.Second, printHello, &wg)

	wg.Add(1)
	go runScheduledTask(5*time.Second, printNumbers, &wg)

	wg.Wait()
}

func runScheduledTask(delay time.Duration, f func(), wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	f()
}

func printHi() {
	println("Hi")
}

func printHello() {
	println("Hello")
}

func printNumbers() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
