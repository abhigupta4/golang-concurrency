package main

import "sync"

type OddEven struct {
	Mutex      *sync.Mutex
	Wg         *sync.WaitGroup
	OddTurn    bool
	IsEvenDone *sync.Cond
	IsOddDone  *sync.Cond
}

func (oe *OddEven) printEven() {
	defer oe.Wg.Done()

	for i := 0; i < 100; i += 2 {
		oe.Mutex.Lock()

		for oe.OddTurn {
			oe.IsOddDone.Wait()
		}
		println(i)
		oe.OddTurn = true
		oe.IsEvenDone.Signal()

		oe.Mutex.Unlock()
	}

}

func (oe *OddEven) printOdd() {
	defer oe.Wg.Done()

	for i := 1; i < 100; i += 2 {
		oe.Mutex.Lock()

		for !oe.OddTurn {
			oe.IsEvenDone.Wait()
		}
		println(i)
		oe.OddTurn = false
		oe.IsOddDone.Signal()

		oe.Mutex.Unlock()
	}
}

func main() {

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	isEvenDone := sync.NewCond(&mutex)
	isOddDone := sync.NewCond(&mutex)
	oe := OddEven{
		Mutex:      &mutex,
		Wg:         &wg,
		IsEvenDone: isEvenDone,
		IsOddDone:  isOddDone,
	}

	wg.Add(2)
	go oe.printEven()
	go oe.printOdd()

	wg.Wait()
}
