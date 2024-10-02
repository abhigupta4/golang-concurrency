package main

import (
	"sync"
	"time"
)

type Slice struct {
	Items []int
	Mutex *sync.RWMutex
}

func (s *Slice) Add(item int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	println("Adding", item)
	s.Items = append(s.Items, item)
}

func (s *Slice) Size() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	size := len(s.Items)
	println("size", size)
	return size
}

func main() {
	s := &Slice{
		Items: make([]int, 0),
		Mutex: &sync.RWMutex{},
	}

	for i := 0; i < 100; i++ {
		go s.Add(i)
	}

	for i := 0; i < 100; i++ {
		go s.Size()
	}

	time.Sleep(1 * time.Second)

	for _, ele := range s.Items {
		println(ele)
	}
}
