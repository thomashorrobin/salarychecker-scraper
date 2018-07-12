package main

import (
	"fmt"
	"strconv"
	"time"
)

type Benchmark struct {
	startTime time.Time
	endTime   time.Time
	pageCount int
}

func (b *Benchmark) start() {
	b.startTime = time.Now()
	b.pageCount = 0
}

func (b *Benchmark) count() {
	b.pageCount++
}

func (b *Benchmark) stop() {
	b.endTime = time.Now()
}

func (b *Benchmark) print() {
	dur := b.endTime.Sub(b.startTime)
	fmt.Println("Time taken: " + dur.String())
	fmt.Println("Number of pages: " + strconv.Itoa(b.pageCount))
}
