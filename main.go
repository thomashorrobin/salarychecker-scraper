package main

import (
	"net/url"
)

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	b := Benchmark{}
	b.Start()
	c := StartCrawl(*url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
		b.IncrementPageCount()
	}
	b.Stop()
	b.Print()
}
