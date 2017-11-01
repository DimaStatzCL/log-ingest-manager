package factory

import (
	"fmt"
	"time"
	"../transport"
)

// Create channel by config
func CreateTransport(out chan <- string) func() {
	return func() {
		transport.CreateWebTransport(out)
	}
}

func CreateNormalizer(in <- chan string, out chan <- string) func() {
	return func() {
		for {
			data := <-in
			out <- "NORMALIZED: " + data
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