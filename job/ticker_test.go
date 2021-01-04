package job

import (
	"testing"
	"time"
)

func Test_JobsWithTicker(t *testing.T) {
	go JobsWithTicker()

	time.Sleep(time.Second * 5)
}
