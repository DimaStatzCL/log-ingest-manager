package factory

import (
	"strconv"
	"time"
	"fmt"
)

// Create channel by config
func CreateTransport(out chan <- string) func() {
	return func() {
		for i := 0; i < 10; i++ {
			out <- "a" + strconv.Itoa(i)
		}
	}
}

func CreateNormalizer(in <- chan string, out chan <- string) func() {
	return func() {
		for {
			data := <-in
			out <- data + "_norm"
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func CreateRoute(in <- chan string) func() {
	return func() {
		for {
			fmt.Println(<-in)
		}
	}
}