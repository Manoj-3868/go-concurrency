package selectstatement

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func MainSelect() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go selectTimeOutWithInCase(&wg)
	wg.Add(1)
	go selectTimeOutOfCase(&wg)
	wg.Wait()
	quit := make(chan string)
	go selectQuitOnChannel(quit)
	time.Sleep(10 * time.Second)
	quit <- "stop"
	fmt.Println(<-quit)
}

func selectTimeOutWithInCase(wg *sync.WaitGroup) {
	defer wg.Done()
	c := boring("Hello : ")
	for {
		select {
		case a := <-c:
			fmt.Println(a)
		case <-time.After(1 * time.Second):
			fmt.Println("you are boring :(")
			return
		}
	}
}

func selectTimeOutOfCase(wg *sync.WaitGroup) {
	defer wg.Done()
	c := boring("Hello : ")
	timeout := time.After(10 * time.Second)
	for {
		select {
		case a := <-c:
			fmt.Println(a)
		case <-timeout:
			fmt.Println("you are boring :(")
			return
		}
	}
}

func selectQuitOnChannel(quit chan string) {
	c := boring("Quit : ")
	for {
		select {
		case a := <-c:
			fmt.Println(a)
		case <-quit:
			quit <- "Bye"
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			if i < 12 {
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return c
}
