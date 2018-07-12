package main

import (
	"net/url"
	"sync"

	page "./page"
)

// PageChannel is an exportable type that allows pages to be pushed to the client
type PageChannel chan page.Page

var urlCheckingChannel chan url.URL
var finishedPageChannel chan page.Page
var checkedURLs map[url.URL]bool
var mux sync.Mutex

// StartCrawl starts a crawl with the initial url as a parameter
func StartCrawl(u url.URL) PageChannel {
	urlCheckingChannel = make(chan url.URL)
	finishedPageChannel = make(chan page.Page)
	checkedURLs = make(map[url.URL]bool)
	mux = sync.Mutex{}

	// start listing to urlCheckingChannel
	go listenOnURLCheckingChannel()

	// put the initial url on the channel
	urlCheckingChannel <- u

	// we return finishedPageChannel to the main function so that is can listen there for parsed pages that have been completed
	return finishedPageChannel
}

func markAsProcessedThenCheckIfCrawlIsDone(url url.URL) {
	mux.Lock()
	checkedURLs[url] = true
	if haveAllQueuedURLsBeenProcessed(checkedURLs) {
		// if all urls are marked as processed then we close the channel to allow the program to exit
		close(finishedPageChannel)
	}
	mux.Unlock()
}

func processURL(url url.URL) {
	page := page.ParseURL(url)
	for _, link := range page.GetLinks() {
		// once a page is returned put all the links on urlCheckingChannel
		urlCheckingChannel <- link
	}
	finishedPageChannel <- page
	// time.Sleep(time.Millisecond * 20)
	go markAsProcessedThenCheckIfCrawlIsDone(url)
}

// this method checks to see if the url is being processed/ has been processed and if it hasn't it'll set off a go routine to process it
func checkThenProcessURL(u url.URL) {
	mux.Lock()
	_, success := checkedURLs[u]
	// if the url isn't in the map then the program hasn't seen it before
	if !success {
		checkedURLs[u] = false
		go processURL(u)
	}
	mux.Unlock()
}

func listenOnURLCheckingChannel() {
	for elem := range urlCheckingChannel {
		go checkThenProcessURL(elem)
	}
}

func haveAllQueuedURLsBeenProcessed(checkedURLs map[url.URL]bool) bool {
	for _, x := range checkedURLs {
		if !x {
			return false
		}
	}
	return true
}
