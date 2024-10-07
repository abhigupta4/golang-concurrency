package main

import (
	"context"
	"sync"
	"time"
)

type ExecutorService struct {
	Wg  *sync.WaitGroup
	Ctx context.Context
}

func (es *ExecutorService) oneTimeWithDelay(task func(), delay time.Duration) {

	es.Wg.Add(1)

	go func() {
		defer es.Wg.Done()
		time.Sleep(delay)

		select {
		case <-es.Ctx.Done():
			return
		default:
			task()
		}
	}()
}

func (es *ExecutorService) scheduledWithDelay(task func(), delay time.Duration, period time.Duration) {
	es.Wg.Add(1)

	go func() {
		defer es.Wg.Done()
		time.Sleep(delay)
		for {
			select {
			case <-es.Ctx.Done():
				return
			default:
				task()
				time.Sleep(period)
			}
		}

	}()
}

func (es *ExecutorService) schedulePeriodWithDelay(task func(), delay time.Duration, period time.Duration) {
	es.Wg.Add(1)

	go func() {
		defer es.Wg.Done()
		time.Sleep(delay)
		ticker := time.NewTicker(period)
		for {
			select {
			case <-es.Ctx.Done():
				return
			case <-ticker.C:
				task()
				time.Sleep(period)
			}
		}

	}()
}

func printHello() {
	println("Hello")
}

func printHi() {
	println("Hi")
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	es := ExecutorService{
		Wg:  &sync.WaitGroup{},
		Ctx: ctx,
	}

	es.oneTimeWithDelay(printHi, time.Second)
	// es.scheduledWithDelay(printHello, time.Second, time.Second)
	es.schedulePeriodWithDelay(printHello, time.Second, time.Second)

	time.Sleep(10 * time.Second)
	cancel()
}
