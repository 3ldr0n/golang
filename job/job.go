package job

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
	fmt.Printf("I am %s, and i'm running.\nI Should be %s\n",
		job.Name, job.Description)
}

func RunJobs() {
	// Channnel with all jobs
	jobs := make(chan Job)
	// Channel to identify when all jobs run
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
		job := SimpleJob{
			Name:        strconv.Itoa(i),
			Description: fmt.Sprintf("This is %d", i),
		}
		jobs <- job
		fmt.Println("Sent job", job.Name)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}
