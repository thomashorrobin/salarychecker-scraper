package main

import "net/url"
import crawler "./crawler"

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	c := crawler.StartCrawl(*url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
	}
}
