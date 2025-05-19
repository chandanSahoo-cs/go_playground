package main

import (
	"sync"
	"time"
	"fmt"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

func main() {
	sc := &safeCounter{
		counts: make(map[string]int),
		mu:     &sync.Mutex{},
	}

	var wg sync.WaitGroup
	key := "foo"

	// Spawn 100 goroutines to increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.inc(key)
		}()
	}

	wg.Wait()

	fmt.Printf("Final value for %s: %d\n", key, sc.val(key))
}

