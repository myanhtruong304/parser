package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID  int
	Job string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d: %s\n", id, task.ID, task.Job)
		// Simulate some work by sleeping for a short duration
		// You can replace this with the actual task processing logic
		// e.g., making an API call, performing a computation, etc.
		// time.Sleep(time.Second)
	}
}

func main() {
	numWorkers := 3
	numTasks := 10

	// Create a task queue
	tasks := make(chan Task)

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start the worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Dispatch tasks to the worker pool
	for i := 1; i <= numTasks; i++ {
		task := Task{
			ID:  i,
			Job: fmt.Sprintf("Job %d", i),
		}
		tasks <- task
	}

	// Close the task channel to indicate that no more tasks will be sent
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks completed")
}
