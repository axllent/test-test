package utils

import (
	"fmt"
	"time"
)

// StopWatch struct
type StopWatch struct {
	StartTime time.Time
}

// StartTimer starts the timer
func StartTimer() StopWatch {
	o := StopWatch{}
	o.StartTime = time.Now()

	return o
}

// Elapsed returns the elapsed time
func (o StopWatch) Elapsed() string {
	t := time.Since(o.StartTime)

	return fmt.Sprintf("%.2f s", t.Seconds())
}
