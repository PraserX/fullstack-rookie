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

type Buffer struct {
	mutex sync.Mutex
	list  []int
}

// Define number of workers
const workerCnt = 5

// Wokers stats
var workerStats = make(chan WorkerStat, workerCnt)

func worker(id int, buffer *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	stat := 0
	for {
		var number int
		buffer.mutex.Lock()
		bufferLen := len(buffer.list)
		if bufferLen > 1 {
			number, buffer.list = buffer.list[0], buffer.list[1:]
		} else if bufferLen == 1 {
			number, buffer.list = buffer.list[0], []int{}
		}
		buffer.mutex.Unlock()

		if number == -1 {
			break
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		// fmt.Printf("[Worker %d]: %d\n", id, number)

		if bufferLen > 0 {
			stat++
		}
	}

	workerStats <- WorkerStat{WorkerID: id, Processed: stat}
}

func main() {
	// Sync package waitgroup
	var wg sync.WaitGroup

	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Create buffer
	buffer := Buffer{}

	go func() {
		// Add 250 random integers to buffer
		for i := 0; i < 250; i++ {
			buffer.mutex.Lock()
			buffer.list = append(buffer.list, rand.Intn(10))
			buffer.mutex.Unlock()
		}
		for i := 0; i < workerCnt; i++ {
			buffer.mutex.Lock()
			buffer.list = append(buffer.list, -1)
			buffer.mutex.Unlock()
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

	close(workerStats)
}
