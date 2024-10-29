package workerpool

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	totaljobs    = 100
	totalworkers = 2
)

func MainWorkerPool() {
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutines running after starting:", runtime.NumGoroutine())

	jobs := make(chan int, totaljobs)
	results := make(chan int, totaljobs)

	for w := 1; w <= totalworkers; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(2 * time.Second)

	for j := range totaljobs + 1 {
		jobs <- j
	}

	for range totaljobs + 1 {
		<-results
	}

	close(jobs)
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("worker ", id, " started")
	var wg sync.WaitGroup
	for data := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Printf("Worker %d started   job %d\n", id, job)

			result := job * 2
			results <- result

			fmt.Printf("Worker %d completed job %d\n", id, job)

		}(data)
	}
	wg.Wait()
}
