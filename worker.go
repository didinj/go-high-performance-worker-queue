// worker.go
package main

import (
	"context"
	"fmt"
)

// Worker represents a single worker that processes jobs
type Worker struct {
	ID       int
	JobQueue <-chan Job      // Read-only channel for jobs
	Context  context.Context // Used for cancellation
}

// Start begins the worker loop in a goroutine
func (w *Worker) Start() {
	go func() {
		fmt.Printf("Worker #%d started\n", w.ID)
		for {
			select {
			case <-w.Context.Done():
				fmt.Printf("Worker #%d stopping...\n", w.ID)
				return
			case job, ok := <-w.JobQueue:
				if !ok {
					fmt.Printf("Worker #%d: job queue closed\n", w.ID)
					return
				}
				job.Process()
			}
		}
	}()
}
