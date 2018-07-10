package main

import "net/url"

func main() {
	url, _ := url.Parse("https://monzo.com")
	page := parseURL(*url)
	page.PrintPage()
}
