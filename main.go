package main

import "net/url"

func main() {
	url, _ := url.Parse("https://monzo.com")
	// page := parseURL(*url)
	// page.PrintPage()
	c := startCrawl(*url)
	// go c.checkURL(*url)
	for elem := range c {
		elem.PrintPage()
	}
}
