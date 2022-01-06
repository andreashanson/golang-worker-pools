package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 100
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)

	for w := 1; w <= 10; w++ {
		go worker(w, jobsChan, completedJobsChan)
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}

	for a := 1; a <= numJobs; a++ {
		<-completedJobsChan
	}
}

func worker(worker int, jobs chan int, completedJobs chan int) {
	for j := range jobs {
		fmt.Println("Worker", worker, "started job", j, "with", len(jobs), "jobs left.")
		time.Sleep(time.Second * 2)
		fmt.Println("Workder", worker, "finished job", j)
		completedJobs <- j
	}
}
