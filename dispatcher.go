// dispatcher.go
package main

import (
	"context"
	"sync"
)

// Dispatcher manages a pool of workers and a job queue
type Dispatcher struct {
	WorkerCount int
	JobQueue    chan Job
	workers     []*Worker
	ctx         context.Context
	cancel      context.CancelFunc
	wg          *sync.WaitGroup
}

// NewDispatcher creates and initializes a Dispatcher
func NewDispatcher(workerCount int, queueSize int) *Dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &Dispatcher{
		WorkerCount: workerCount,
		JobQueue:    make(chan Job, queueSize),
		ctx:         ctx,
		cancel:      cancel,
		wg:          &sync.WaitGroup{},
	}
}

// Start initializes the worker pool and begins processing jobs
func (d *Dispatcher) Start() {
	for i := 1; i <= d.WorkerCount; i++ {
		worker := &Worker{
			ID:       i,
			JobQueue: d.JobQueue,
			Context:  d.ctx,
		}
		d.workers = append(d.workers, worker)
		worker.Start()
	}
}

// Submit enqueues a job into the job queue
func (d *Dispatcher) Submit(job Job) {
	d.wg.Add(1)
	go func() {
		d.JobQueue <- job
		d.wg.Done()
	}()
}

// Stop gracefully shuts down the dispatcher and its workers
func (d *Dispatcher) Stop() {
	// Cancel context to stop all workers
	d.cancel()

	// Wait for any pending submissions to complete
	d.wg.Wait()

	// Close the job queue
	close(d.JobQueue)
}
