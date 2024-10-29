package queue

import (
	"fmt"
	"sync"
	"time"
)

const (
	QueueLength  = 5
	ProcessCount = 100
)

func MainQueue() {
	queue := make(chan struct{}, QueueLength)
	wg := sync.WaitGroup{}
	wg.Add(ProcessCount)
	for i := range ProcessCount {
		process(i, queue, &wg)
	}
	wg.Wait()
}

func process(id int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}
	go func() {
		defer wg.Done()
		fmt.Println("Processed : ", id)
		time.Sleep(1 * time.Second)
		<-queue
	}()
}
