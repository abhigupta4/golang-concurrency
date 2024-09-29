package main

import "fmt"

func unbufferedChannel() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Println(value)
}

func bufferedChannel() {
	buffCh := make(chan int, 10)
	go addNumbers(buffCh)

	for i := 0; i < 5; i++ {
		fmt.Println(<-buffCh)
	}
}

func addNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	// unbufferedChannel()

	bufferedChannel()
}
