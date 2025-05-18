package main

import (
	"time"
	"fmt"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for{
	select {
	case <-snapshotTicker :
		takeSnapshot(logChan)
	case <-saveAfter :
		saveSnapshot(logChan)
		return
	default :
	waitForData(logChan)
	time.Sleep(500*time.Millisecond)
	}
}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	logChan := make(chan string)

	// Create two ticker channels for snapshotTicker and saveAfter
	snapshotTicker := time.Tick(1 * time.Second)   // ticks every 1 second
	saveAfter := time.After(5 * time.Second)        // fires once after 5 seconds

	go saveBackups(snapshotTicker, saveAfter, logChan)

	// Read from logChan until it is closed
	for msg := range logChan {
		fmt.Println(msg)
	}

	fmt.Println("Backup process completed.")
}