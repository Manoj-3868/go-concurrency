package fanin

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func MainFanIn() {
	/*
		outputChan := fanIn(boring("hello"), boring("world"))
		for range 20 {
			fmt.Println(<-outputChan)
		}
	*/
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(chan1 <-chan string, chan2 <-chan string) (output chan string) {
	output = make(chan string)
	go func() {
		for {
			output <- <-chan1
		}
	}()
	go func() {
		for {
			output <- <-chan2
		}
	}()
	return
}

func fanInUsingSelect(chan1 <-chan string, chan2 <-chan string) (output chan string) {
	output = make(chan string)
	go func() {
		for {
			select {
			case m1 := <-chan1:
				output <- m1
			case m2 := <-chan2:
				output <- m2
			}
		}
	}()
	return
}

//ASynchronous
//Daemon
//Queue
//pooled goRoutine -> wait group
//Bounded concuurency -> wait group + channels
//Genarator pattern
//Signal pattern -> channel with empty struct
//Timeout pattern ->
//FanIn
//Fanout
//Pipeline -> connecting multiple cahnnels
//CSP ()
//Pub/Sub

//convert ram to cpu , cpu to ram , ram to storage
