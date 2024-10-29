package pubsubv2

import (
	"fmt"
	"sync"
	"testing"
)

var mu = sync.RWMutex{}

const (
	topic1     = "hello"
	topic1Data = "data1"
	topic2     = "world"
	topic2Data = "data2"
	datasize   = 100
)

func TestMain(t *testing.T) {
	results := map[string][]string{}

	pubsub := New()
	pubsub.NewTopic(topic1)
	pubsub.NewTopic(topic2)

	exit1 := make(chan struct{})
	exit2 := make(chan struct{})

	exit := func() {
		exit1 <- struct{}{}
		exit2 <- struct{}{}
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		defer exit()
		for i := range datasize {
			err := pubsub.Publish(topic1, topic1Data+fmt.Sprintf(" %d", i))
			// err := pubsub.Publish(topic1, topic1Data)
			if err != nil {
				t.Fail()
				t.Log(err)
			}
			err = pubsub.Publish(topic2, topic2Data+fmt.Sprintf(" %d", i))
			// err = pubsub.Publish(topic2, topic2Data)
			if err != nil {
				t.Fail()
				t.Log(err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		sub1, err := pubsub.Subscribe(topic1)
		if err != nil {
			t.Fail()
			t.Log(err)
		}
		for {
			select {
			case data := <-sub1:
				mu.Lock()
				results[topic1] = append(results[topic1], data)
				mu.Unlock()
			case <-exit1:
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		sub2, err := pubsub.Subscribe(topic2)
		if err != nil {
			t.Fail()
			t.Log(err)
		}
		for {
			select {
			case data := <-sub2:
				mu.Lock()
				results[topic2] = append(results[topic2], data)
				mu.Unlock()
			case <-exit2:
				return
			}
		}
	}()
	wg.Wait()

	_, ok := results[topic1]
	if !ok {
		t.FailNow()
		t.Log("ok : ", ok)
	}
	_, ok = results[topic2]
	if !ok {
		t.FailNow()
		t.Log("ok : ", ok)
	}

	if len(results[topic1]) != datasize {
		t.FailNow()
		t.Log("topic1")
	}

	if len(results[topic2]) != datasize {
		t.FailNow()
		t.Log("topic1")
	}

	for i := range results[topic1] {
		if results[topic1][i] != topic1Data+fmt.Sprintf(" %d", i) {
			t.FailNow()
		}
	}
	for i := range results[topic2] {
		if results[topic2][i] != topic2Data+fmt.Sprintf(" %d", i) {
			t.FailNow()
		}
	}
}
