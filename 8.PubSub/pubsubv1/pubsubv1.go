// simple pub/sub
package pubsubv1

type PubSub struct {
	subscribers []chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make([]chan string, 0),
	}
}

func (pubsub *PubSub) NewSub() <-chan string {
	sub := make(chan string)
	pubsub.subscribers = append(pubsub.subscribers, sub)
	return sub
}

func (pubsub *PubSub) Publish(message string) {
	for i := range pubsub.subscribers {
		pubsub.subscribers[i] <- message
	}
}
