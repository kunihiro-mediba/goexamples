package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	MaxThread = 10
	Tasks = 100
)

func sub(threads chan bool, i int) {
	log.Printf("[%2d] start", i)
	ms := rand.Int() % 1000 + 500 // 500ms 〜 1500ms
	time.Sleep(time.Duration(ms) * time.Millisecond)
	log.Printf("[%2d] end", i)
	<-threads
}

func main() {
	startTime := time.Now()

	threads := make(chan bool, MaxThread)
	for i := 0; i < Tasks; i++ {
		threads <- true
		go sub(threads, i)
	}

	// すべてのチャンネルが空くのを待つ
	for i := 0; i < MaxThread; i++ {
		threads <- false
	}

	endTime := time.Now()
	durationMs := endTime.Sub(startTime).Milliseconds()

	log.Printf("finish %d ms", durationMs)
}
