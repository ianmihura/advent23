package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Root struct {
	n int
	s float64
}

func worker(id int, jobs <-chan int, results chan<- Root) {
	fmt.Println("Im worker", id)
	for j := range jobs { // pulling job from jobs channel
		time.Sleep(time.Second)
		results <- Root{n: j, s: math.Sqrt(float64(j))}
		fmt.Println("Worker", id, "finished job", j)
	}
}

func workerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)     // buffered channel
	results := make(chan Root, numJobs) // buffered channel

	// Create workers
	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results) // spinup workers
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j // sending jobs via channel
	}

	// Receive results
	for j := 1; j <= numJobs; j++ {
		res := <-results // pulling result from channel
		fmt.Println(res.n, res.s)
	}

	close(jobs)    // gracefully closes channels
	close(results) // gracefully closes channels

	/*
		Note on closed channels:
		Sending to a closed channel will panic()
		Trying to close a closed channel will panic()
		Receiving from a closed channel will succeede with a zero or nil value
	*/
}

func someSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		mess := "in func 1"
		c1 <- mess
		stress := <-c2 // normally blocked by func2
		fmt.Println("from func1:", mess, stress)
	}()
	go func() {
		mess := "in func 2"
		var stress string
		select {
		case c2 <- mess:
			fmt.Println("c2 <- mess")
		case stress = <-c1:
			fmt.Println("stress = <-c1")
		default: // question: why does this execute instead of case?
			fmt.Println("DONE")
		}
		fmt.Println("from func2:", mess, stress)
		c2 <- mess // send info to func1
	}()

	// allow the goroutines to execute before main() finishes
	time.Sleep(time.Second)
}

func addSum() float64 {
	const N = 1_000_000
	var data [N]float64
	var sum float64
	for i := range N {
		data[i] = 10.0 * rand.Float64()
	}
	for _, d := range data {
		sum = sum + d
	}
	return sum / N
}

func waitgroup() {
	var wg sync.WaitGroup
	threads := 50
	wg.Add(threads)

	start := time.Now().UnixMilli()
	for i := range threads {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "sum is", addSum())
		}(i)
	}

	wg.Wait()
	fmt.Println("DONE in ms", time.Now().UnixMilli()-start)
}

func main() {
	// workerPool()
	// someSelect()
	// waitgroup()
	PhiProblem()
}
