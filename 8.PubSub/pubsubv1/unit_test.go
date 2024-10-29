package pubsubv1

import (
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	pub := NewPubSub()

	testData := "hello"
	testDataSize := 100

	sub1Results := []string{}
	sub2Results := []string{}

	sub1 := pub.NewSub()
	sub2 := pub.NewSub()

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
		time.Sleep(2 * time.Second)
		for range testDataSize {
			pub.Publish(testData)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case result := <-sub1:
				sub1Results = append(sub1Results, result)
			case <-exit1:
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case result := <-sub2:
				sub2Results = append(sub2Results, result)
			case <-exit2:
				return
			}
		}
	}()

	wg.Wait()

	if len(sub1Results) != testDataSize {
		t.Fail()
	}

	if len(sub2Results) != testDataSize {
		t.Fail()
	}

	for i := range sub1Results {
		if sub1Results[i] != testData {
			t.Fail()
		}
	}

	for i := range sub2Results {
		if sub2Results[i] != testData {
			t.Fail()
		}
	}

	t.Log("Test Successfull")

}
