package main

import "fmt"
import "./logic"

func main () {
	fmt.Println("initialize channels")

	// TODO: load configuration and use logging
	maxChannelSize := 100000
	degreeOfParallelism := 0

	// Start pipeline with degree of parallelism
	logic.StartPipeline(maxChannelSize, degreeOfParallelism)
	fmt.Printf("pipeline started: %d, %d", maxChannelSize, degreeOfParallelism)
	fmt.Scanln()

	// Stop pipeline
	logic.StopPipeline()
	fmt.Println("pipeline stopped")
}


