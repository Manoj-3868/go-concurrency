package fanout

import "fmt"

func MainFanOut() {
	values := []int{22, 44, 55, 66, 77, 88}
	input := generator(values)

	out1 := fanOut(input)
	out2 := fanOut(input)
	out3 := fanOut(input)
	out4 := fanOut(input)

	for range values {
		select {
		case i := <-out1:
			fmt.Println("Goroutine 1 Processed : ", i)
		case i := <-out2:
			fmt.Println("Goroutine 2 Processed : ", i)
		case i := <-out3:
			fmt.Println("Goroutine 3 Processed : ", i)
		case i := <-out4:
			fmt.Println("Goroutine 4 Processed : ", i)
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range in {
			out <- data
		}
	}()
	return out
}

func generator(work []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range work {
			out <- work[i]
		}
	}()
	return out
}
