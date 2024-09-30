package main

import "time"

type Queue struct {
	Items chan int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		Items: make(chan int, capacity),
	}
}

func main() {
	q := NewQueue(5)

	for i := 0; i < 10; i++ {
		go produce(i, q)
	}

	for i := 0; i < 1; i++ {
		go consume(i, q)
	}

	time.Sleep(2 * time.Second)
}

func produce(id int, q *Queue) {
	start := id * 100
	for i := start; i < start+10; i++ {
		q.Items <- i
		println("Producer", id, "Item", i)
	}
}

func consume(id int, q *Queue) {
	for i := range q.Items {
		println("Consumer", id, "Item", i)
	}
}
