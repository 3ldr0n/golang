package job

import (
	"fmt"
	"time"
)

func JobsWithTicker() {
	// Creates a ticker for a scheduled job.
	jobs := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-jobs.C:
			fmt.Println("Running scheduled job")
		}
	}
}
