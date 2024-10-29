package sequentialchannel

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type Message struct {
	str  string
	wait chan bool
}

func MainSequential() {
	ch := make(chan Message)
	for i := range 5 {
		go speaker(ch, fmt.Sprintf("Speaker  : %d", i))
	}
	for range 5 {
		msg1 := <-ch
		fmt.Println(msg1.str)
		msg2 := <-ch
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func speaker(ch chan Message, name string) {
	for i := range 3 {
		waitForit := make(chan bool)
		msg := Message{
			str:  fmt.Sprintf("%s : %d", name, i),
			wait: waitForit,
		}
		ch <- msg
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		<-waitForit
	}
}
