package pipeline

import "fmt"

func MainPipeline() {
	values := []int{22, 44, 55, 66, 77, 88}
	input := generator(values)
	out := filter(input)
	out = square(out)
	out = half(out)

	for data := range out {
		fmt.Println("DATA : ", data)
	}
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		defer fmt.Println("defer filter...")
		for data := range in {
			if data%2 == 0 {
				out <- data
			}
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		defer fmt.Println("defer square...")
		for data := range in {
			out <- data * data
		}
	}()
	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		defer fmt.Println("defer half...")
		for data := range in {
			out <- data / 2
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
