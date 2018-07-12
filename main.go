package main

import (
	"net/url"
)

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	b := Benchmark{}
	b.start()
	c := StartCrawl(*url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
		b.count()
	}
	b.stop()
	b.print()
}
