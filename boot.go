package main

import "fmt"
import "./factory"


func main () {
	fmt.Println("initialize channels")
	raw_data := make(chan string, 1000)
	normalized_data := make(chan string, 1000)

	route := factory.CreateRoute(normalized_data)
	transport := factory.CreateTransport(raw_data)
	normalizer := factory.CreateNormalizer(raw_data, normalized_data)

	go route()
	go transport()
	go normalizer()

	fmt.Println("pipeline started")
	fmt.Scanln()
}
