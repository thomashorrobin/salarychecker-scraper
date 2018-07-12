package main

import (
	"fmt"
	"strconv"
	"time"
)

type benchmark struct {
	startTime time.Time
	endTime   time.Time
	pageCount int
}

func (b *benchmark) start() {
	b.startTime = time.Now()
	b.pageCount = 0
}

func (b *benchmark) incrementPageCount() {
	b.pageCount++
}

func (b *benchmark) stop() {
	b.endTime = time.Now()
}

func (b *benchmark) print() {
	dur := b.endTime.Sub(b.startTime)
	fmt.Println("Time taken: " + dur.String())
	fmt.Println("Number of pages: " + strconv.Itoa(b.pageCount))
}
