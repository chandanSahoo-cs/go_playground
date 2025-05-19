package main

import (
	"sync"
	"time"
	"fmt"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}


func main() {
	sc := &safeCounter{
		counts: make(map[string]int),
		mu:     &sync.RWMutex{},
	}

	var wg sync.WaitGroup
	key := "foo"

	// Spawn 100 goroutines to increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sc.inc(key)
			wg.Done()
		}()
	}

	// Wait for all increments to finish
	wg.Wait()

	// Read the final value
	fmt.Printf("Final value for key '%s': %d\n", key, sc.val(key))
}