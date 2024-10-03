package main

import (
	"sync"
	"time"
)

type Philosopher struct {
	Id        int
	LeftFork  *sync.Mutex
	RightFork *sync.Mutex
}

func (p *Philosopher) live(wg *sync.WaitGroup, semaphore chan int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		println("Philosopher", p.Id, "done thinking")

		semaphore <- 1
		p.LeftFork.Lock()
		p.RightFork.Lock()
		println("Philosopher", p.Id, "done eating")
		p.RightFork.Unlock()
		p.LeftFork.Unlock()
		<-semaphore
	}
}

func main() {
	pCount := 5
	forks := make([]*sync.Mutex, pCount)
	wg := sync.WaitGroup{}
	semaphore := make(chan int, pCount-1)

	for i := 0; i < pCount; i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < pCount; i++ {
		wg.Add(1)
		p := Philosopher{
			Id:        i + 1,
			LeftFork:  forks[i],
			RightFork: forks[(i+1)%pCount],
		}

		go p.live(&wg, semaphore)
	}

	wg.Wait()
}
