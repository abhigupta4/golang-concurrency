package main

import "time"

func work(id int, jobs chan int) {

	for job := range jobs {
		println("worker", id, "started job", job)
		println("worker", id, "finished job", job)
	}

}

func main() {

	jobs := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go work(i, jobs)
	}

	for j := 0; j < 100; j++ {
		jobs <- j
	}

	println("waiting for jobs to finish")
	time.Sleep(2 * time.Second)
	close(jobs)
}
