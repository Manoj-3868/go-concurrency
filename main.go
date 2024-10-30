package main

import (
	"fmt"
	"time"

	workerpool "server/6.WorkerPool"

	"golang.org/x/exp/rand"
)

func main() {
	// fanin.MainFanIn()
	// sequentialchannel.MainSequential()
	// selectstatement.MainSelect()
	// fanout.MainFanOut()
	// pipeline.MainPipeline()
	workerpool.MainWorkerPool()
	// queue.MainQueue()
	// pubsub.MainPubSub()
}

func initialBoring() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
