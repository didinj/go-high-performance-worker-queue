// job.go
package main

import (
	"fmt"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID      int
	Message string
}

// Process executes the actual job logic
// func (j Job) Process() {
// 	fmt.Printf("Processing Job #%d: %s\n", j.ID, j.Message)
// 	time.Sleep(1 * time.Second) // simulate workload
// 	fmt.Printf("Finished Job #%d\n", j.ID)
// }

// job.go (modified)
func (j Job) Process() {
	fmt.Printf("Processing Job #%d: %s\n", j.ID, j.Message)
	time.Sleep(1 * time.Second)

	// Simulate error
	if j.ID%7 == 0 {
		fmt.Printf("⚠️ Job #%d failed: simulated error\n", j.ID)
		return
	}

	fmt.Printf("Finished Job #%d\n", j.ID)
}
