package main

import (
	"context"
	"sync"
)

type Queue struct {
	Items      []int
	Capacity   int
	Mutex      *sync.Mutex
	IsNotFull  *sync.Cond
	IsNotEmpty *sync.Cond
}

func (q *Queue) Add(id int, val int) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for len(q.Items) == q.Capacity {
		q.IsNotFull.Wait()
	}

	println("Produced ", val, " By ", id)
	q.Items = append(q.Items, val)
	q.IsNotEmpty.Signal()
}

func (q *Queue) Get(id int, context context.Context) (int, bool) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for len(q.Items) == 0 {
		select {
		case <-context.Done():
			return 0, true
		default:
			q.IsNotEmpty.Wait()
		}
	}

	first := q.Items[0]
	q.Items = q.Items[1:]

	println("Consumed ", first, " By", id)
	q.IsNotFull.Signal()
	return first, false
}

func NewQueue(capacity int) *Queue {
	mutex := sync.Mutex{}

	return &Queue{
		Items:      make([]int, 0),
		Capacity:   capacity,
		Mutex:      &mutex,
		IsNotFull:  sync.NewCond(&mutex),
		IsNotEmpty: sync.NewCond(&mutex),
	}
}

func main() {
	q := NewQueue(5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go produce(i, q, wg)
	}

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		go consume(i, q, ctx)
	}

	wg.Wait()
	cancel()
}

func produce(id int, q *Queue, wg *sync.WaitGroup) {

	defer wg.Done()
	offset := id * 100
	for i := offset; i < offset+5; i++ {
		q.Add(id, i)
	}
}

func consume(id int, q *Queue, context context.Context) {
	for {
		_, closed := q.Get(id, context)
		if closed {
			return
		}
	}
}
