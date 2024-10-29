// topic based pub/sub
package pubsubv2

import (
	"errors"
	"sync"
)

var (
	ErrInvalidTopic = errors.New("")
)

type PubSub struct {
	mu            sync.Mutex
	subscriptions map[string][]chan string
}

func New() (pubsub *PubSub) {
	return &PubSub{
		mu:            sync.Mutex{},
		subscriptions: make(map[string][]chan string),
	}
}

func (model *PubSub) NewTopic(name string) {
	model.mu.Lock()
	defer model.mu.Unlock()
	_, ok := model.subscriptions[name]
	if ok {
		return
	}
	model.subscriptions[name] = make([]chan string, 0)
}

func (model *PubSub) Subscribe(topic string) (subcription <-chan string, err error) {
	model.mu.Lock()
	defer model.mu.Unlock()
	_, ok := model.subscriptions[topic]
	if !ok {
		err = ErrInvalidTopic
		return
	}
	sub := make(chan string)
	model.subscriptions[topic] = append(model.subscriptions[topic], sub)
	subcription = sub
	return
}

func (model *PubSub) Publish(topic string, msg string) (err error) {
	model.mu.Lock()
	defer model.mu.Unlock()
	_, ok := model.subscriptions[topic]
	if !ok {
		err = ErrInvalidTopic
		return
	}
	for _, subcription := range model.subscriptions[topic] {
		subcription <- msg
	}
	return
}

func (model *PubSub) UnSubscribe(topic string, subcription <-chan string) (err error) {
	model.mu.Lock()
	defer model.mu.Unlock()
	_, ok := model.subscriptions[topic]
	if !ok {
		err = ErrInvalidTopic
		return
	}
	for i, sub := range model.subscriptions[topic] {
		if subcription == sub {
			model.subscriptions[topic] = append(model.subscriptions[topic][:i], model.subscriptions[topic][i+1:]...)
		}
	}
	return
}
