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
		urlDefault, _ := url.Parse("https://www.totaljobs.com")
		url = *urlDefault
	}
	b := benchmark{}
	b.start()
	c := StartCrawl(url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
		b.incrementPageCount()
	}
	b.stop()
	b.print()
}
