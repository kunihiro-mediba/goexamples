package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	MaxTPS = 10
	MaxThread = 30
	Tasks = 300
	IntervalMS = 1000 / MaxTPS
)

func sub(threads chan bool, i int) {
	log.Printf("[%3d] start", i)
	ms := rand.Int() % 2900 + 100 // 100ms 〜 3000ms
	time.Sleep(time.Duration(ms) * time.Millisecond)
	log.Printf("[%3d] end (%d ms)", i, ms)
	<-threads
}

func main() {
	threads := make(chan bool, MaxThread)

	startTime := time.Now()
	for i := 0; i < Tasks; i++ {
		threads <- true
		tti := time.Now().Sub(startTime).Milliseconds()
		border := int64(IntervalMS * i)
		if tti < border {
			time.Sleep(time.Duration(border - tti) * time.Millisecond)
		}
		go sub(threads, i)
	}

	// すべてのチャンネルが空くのを待つ
	for i := 0; i < MaxThread; i++ {
		threads <- false
	}

	endTime := time.Now()
	durationMs := endTime.Sub(startTime).Milliseconds()
	tps := float64(Tasks) / float64(durationMs / 1000)

	log.Printf("finish %d ms (%.2f tps)", durationMs, tps)
}
