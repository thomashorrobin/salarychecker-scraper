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

func (b *Benchmark) Start() {
	b.startTime = time.Now()
	b.pageCount = 0
}

func (b *Benchmark) IncrementPageCount() {
	b.pageCount++
}

func (b *Benchmark) Stop() {
	b.endTime = time.Now()
}

func (b *Benchmark) Print() {
	dur := b.endTime.Sub(b.startTime)
	fmt.Println("Time taken: " + dur.String())
	fmt.Println("Number of pages: " + strconv.Itoa(b.pageCount))
}
