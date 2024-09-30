package main

import (
	"sync"
	"time"
)

type Queue struct {
	mutex    *sync.Mutex
	items    []int
	capacity int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		mutex:    &sync.Mutex{},
		items:    make([]int, 0),
		capacity: capacity,
	}
}

func (q *Queue) Add(data int) {
	for {
		q.mutex.Lock()

		if len(q.items) < q.capacity {
			q.items = append(q.items, data)
			println("produced ", data)
			q.mutex.Unlock()
			return
		}
		q.mutex.Unlock()
	}
}

func (q *Queue) Remove() int {
	for {
		q.mutex.Lock()

		if len(q.items) > 0 {
			first := q.items[0]
			q.items = q.items[1:]
			println("consumed ", first)
			q.mutex.Unlock()
			return first
		}

		q.mutex.Unlock()
	}
}

func main() {
	q := NewQueue(30)

	go startProducer(q)
	go startConsumer(q)
	go startConsumer(q)
	go startConsumer(q)
	go startConsumer(q)

	time.Sleep(100 * time.Second)
}

func startProducer(q *Queue) {
	for i := 0; i < 1000; i++ {
		q.Add(i)
	}
}

func startConsumer(q *Queue) {
	for {
		_ = q.Remove()
	}
}
