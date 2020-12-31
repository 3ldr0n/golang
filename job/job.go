package main

import (
	"fmt"
	"strconv"
)

type Job interface {
	ID() string
	Run()
}

type SimpleJob struct {
	Name        string
	Description string
}

func (job SimpleJob) ID() string {
	return job.Name
}

func (job SimpleJob) Run() {
	fmt.Printf("I am %s, and i'm running.\n", job.Name)
}

func main() {
	jobs := make(chan Job)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				job.Run()
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- SimpleJob{
			Name:        strconv.Itoa(i),
			Description: fmt.Sprintf("This is %d", i),
		}
		fmt.Println("Sent job", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}
