package main

import (
	"net/url"

	crawler "./crawler"
)

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	b := Benchmark{}
	b.start()
	c := crawler.StartCrawl(*url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
		b.count()
	}
	b.stop()
	b.print()
}
