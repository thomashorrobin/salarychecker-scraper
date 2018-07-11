package main

import "net/url"

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	c := initCrawler()
	go c.checkURL(*url)
	for elem := range c.finishedPageChannel {
		elem.PrintPage()
	}
}
