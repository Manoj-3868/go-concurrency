package pubsub

import (
	"fmt"
	"time"
)

type PubSub struct {
	subscribers []chan string
	// confirm     chan struct{}
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make([]chan string, 0),
		// confirm:     make(chan struct{}),
	}
}

func (pubsub *PubSub) NewSub() <-chan string {
	sub := make(chan string)
	pubsub.subscribers = append(pubsub.subscribers, sub)
	return sub
}

func (pubsub *PubSub) Publish(message string) {
	// go func() {
	// defer pubsub.Start()
	// pubsub.Stop()
	for i := range pubsub.subscribers {
		pubsub.subscribers[i] <- message
	}
	// }()

}

// func (pubsub *PubSub) Start() {
// 	go func() {
// 		pubsub.confirm <- struct{}{}
// 	}()
// }
// func (pubsub *PubSub) Stop() {
// 	go func() {
// 		<-pubsub.confirm
// 	}()
// }

func (pubsub *PubSub) PublishToActiveSubs(message string) {
	for i := range pubsub.subscribers {
		select {
		case pubsub.subscribers[i] <- message:
		default:
			continue
		}
	}
}

func MainPubSub() {
	pub := NewPubSub()
	hello := pub.NewSub()
	world := pub.NewSub()
	go func() {
		time.Sleep(2 * time.Second)
		// pub.Start()
		for i := 0; i < 10; i++ {
			pub.Publish(fmt.Sprintf("%d", i))
		}
		// pub.Stop()
		fmt.Println("loop exited")
	}()
	go func() {
		for msg := range hello {
			fmt.Println(msg, "to subcriber 0")
		}
	}()
	go func() {
		for msg := range world {
			fmt.Println(msg, "to subcriber 1")
		}
	}()
	go func() {
		sub3 := pub.NewSub()
		// time.Sleep(10 * time.Second)
		for msg := range sub3 {
			fmt.Println(msg, "to subcriber 2")
		}
	}()
	time.Sleep(10 * time.Second)
}
