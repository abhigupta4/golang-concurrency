package main

import "fmt"

func printNumber(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 100; i++ {
		go printNumber(i)
	}
}
