package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. GENERATOR: Converts a slice of data into a readable channel.
func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // Always close the channel when done!
	}()
	return out
}

// 2. WORKER (Fan-out): Reads from the input channel, processes data, and sends to an output channel.
func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			// Simulating an expensive operation (e.g., API call, DB query, heavy math)
			time.Sleep(100 * time.Millisecond) 
			fmt.Printf("[Worker %d] Processed: %d\n", id, n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 3. MERGE (Fan-in): Combines multiple channels into a single output channel.
func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	// Helper function to forward data from a single channel to the merged channel
	output := func(c <-chan int) {
		for n := range c {
			multiplexedStream <- n
		}
		wg.Done()
	}

	// Increment WaitGroup by the number of channels we are merging
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Start a goroutine to close the multiplexed channel once all workers are done
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func main() {
	start := time.Now()

	// Initial data source
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputChan := generator(inputs)

	// --- FAN OUT ---
	// Distribute work across 3 concurrent worker goroutines
	numWorkers := 3
	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = worker(i+1, inputChan)
	}

	// --- FAN IN ---
	// Merge all worker output channels back into a single result channel
	finalResultChan := fanIn(workers...)

	// Consume the aggregated results
	fmt.Println("--- Pipeline Started ---")
	for result := range finalResultChan {
		fmt.Printf("Final Result Received: %d\n", result)
	}

	fmt.Printf("\nTotal Execution Time: %v\n", time.Since(start))
}