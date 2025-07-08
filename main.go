// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// func main() {
// 	workerCount := 4
// 	queueSize := 10

// 	fmt.Println("Starting dispatcher...")
// 	dispatcher := NewDispatcher(workerCount, queueSize)
// 	dispatcher.Start()

// 	// Simulate job production
// 	go func() {
// 		for i := 1; i <= 20; i++ {
// 			job := Job{
// 				ID:      i,
// 				Message: fmt.Sprintf("Job number %d", i),
// 			}
// 			fmt.Printf("Submitting Job #%d\n", job.ID)
// 			dispatcher.Submit(job)
// 			time.Sleep(500 * time.Millisecond) // simulate incoming rate
// 		}
// 	}()

// 	// Handle OS signals for graceful shutdown
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

// 	<-sigChan // wait for interrupt
// 	fmt.Println("\nShutting down gracefully...")
// 	dispatcher.Stop()
// 	fmt.Println("All workers stopped. Exiting.")
// }

// main.go (partial)

func main() {
	workerCount := 4
	queueSize := 10
	totalJobs := 20

	fmt.Println("Starting dispatcher...")
	dispatcher := NewDispatcher(workerCount, queueSize)
	dispatcher.Start()

	start := time.Now()

	// Simulate job production
	go func() {
		for i := 1; i <= totalJobs; i++ {
			job := Job{
				ID:      i,
				Message: fmt.Sprintf("Job number %d", i),
			}
			fmt.Printf("Submitting Job #%d\n", job.ID)
			dispatcher.Submit(job)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// OS signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nShutting down gracefully...")
	dispatcher.Stop()
	elapsed := time.Since(start)
	fmt.Printf("Processed %d jobs in %s\n", totalJobs, elapsed)
	fmt.Println("All workers stopped. Exiting.")
}
