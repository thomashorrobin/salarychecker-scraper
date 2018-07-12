package main

import (
	"net/url"
	"os"
)

func main() {
	var url url.URL
	if len(os.Args) >= 2 {
		urlFlag, _ := url.Parse(os.Args[1])
		url = *urlFlag
	} else {
		urlDefault, _ := url.Parse("https://monzo.com")
		url = *urlDefault
	}
	b := Benchmark{}
	b.Start()
	c := StartCrawl(url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
		b.IncrementPageCount()
	}
	b.Stop()
	b.Print()
}
