package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkerStat struct {
	WorkerID  int
	Processed int
}

// Define number of workers
const workerCnt = 5

// Wokers stats
var workerStats = make(chan WorkerStat, workerCnt)

func worker(id int, buffer *chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	stat := 0
	for number := range *buffer {
		if number == -1 {
			break
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		// fmt.Printf("[Worker %d]: %d\n", id, number)
		stat++
	}

	workerStats <- WorkerStat{WorkerID: id, Processed: stat}
}

func main() {
	// Sync package waitgroup
	var wg sync.WaitGroup

	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Create buffer
	buffer := make(chan int, 25)

	go func() {
		// Add 250 random integers to buffer
		for i := 0; i < 250; i++ {
			buffer <- rand.Intn(10)
		}
		for i := 0; i < workerCnt; i++ {
			buffer <- -1
		}
	}()

	for i := 1; i <= workerCnt; i++ {
		wg.Add(1)                  // Add 1 worker to worker counter
		go worker(i, &buffer, &wg) // Start worker
	}

	// Wait for all workers
	wg.Wait()

	sum := 0
	for i := 0; i < workerCnt; i++ {
		stat := <-workerStats
		fmt.Printf("Worker [%d]: %d\n", stat.WorkerID, stat.Processed)
		sum += stat.Processed
	}

	fmt.Println("SUM:", sum)

	close(buffer)
	close(workerStats)
}
