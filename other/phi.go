package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	PHILOSOPHERS = 5
	FORKS        = 5
	SERVINGS     = 3
	SCOOP        = 1
)

type Philosopher struct {
	id    int
	lFork *sync.Mutex
	rFork *sync.Mutex
	scoop *sync.Mutex
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for range SERVINGS {
		think := rand.Intn(5)
		fmt.Println(p.id, "thinking for", think)
		time.Sleep(time.Duration(think) * time.Second)

		p.scoop.Lock()
		fmt.Println(p.id, "took scoop")
		time.Sleep(time.Second)
		p.scoop.Unlock()
		fmt.Println(p.id, "dropped scoop")

		p.lFork.Lock()
		fmt.Println(p.id, "took left")
		p.rFork.Lock()
		fmt.Println(p.id, "took right")

		eat := rand.Intn(5)
		fmt.Println(p.id, "eating for", eat)
		time.Sleep(time.Duration(eat) * time.Second)

		p.lFork.Unlock()
		fmt.Println(p.id, "dropped left")
		p.rFork.Unlock()
		fmt.Println(p.id, "dropped right")
	}
	fmt.Println(p.id, "finished eating")
}

func PhiProblem() {
	forks := make([]*sync.Mutex, FORKS)
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}
	scoop := sync.Mutex{}

	philosophers := make([]*Philosopher, PHILOSOPHERS)
	for i := range philosophers {
		philosophers[i] = &Philosopher{
			id:    i,
			lFork: forks[i],
			rFork: forks[(i+1)%FORKS],
			scoop: &scoop,
		}
	}

	var wg sync.WaitGroup
	wg.Add(PHILOSOPHERS)

	for _, philosopher := range philosophers {
		go philosopher.Eat(&wg)
	}

	wg.Wait()
}
