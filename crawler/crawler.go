package crawler

import (
	"net/url"
	"sync"

	"../page"
)

// this represents a channel for passing parsed pages to a ui or other enduser
type PageChannel chan page.Page

var urlCheckingChannel chan url.URL
var finishedPageChannel chan page.Page
var checkedURLs map[url.URL]bool
var mux sync.Mutex

// xxx
func StartCrawl(u url.URL) PageChannel {
	urlCheckingChannel = make(chan url.URL)
	finishedPageChannel = make(chan page.Page)
	checkedURLs = make(map[url.URL]bool)
	mux = sync.Mutex{}
	go listenOnURLCheckingChannelLimited(5)
	urlCheckingChannel <- u
	return finishedPageChannel
}

func markAsProcessedThenCheckIfCrawlIsDone(url url.URL) {
	mux.Lock()
	checkedURLs[url] = true
	if haveAllQueuedURLsBeenProcessed(checkedURLs) {
		close(finishedPageChannel)
	}
	mux.Unlock()
}

func processURL(url url.URL) {
	page := page.ParseURL(url)
	for _, link := range page.GetLinks() {
		urlCheckingChannel <- link
	}
	finishedPageChannel <- page
	// time.Sleep(time.Millisecond * 20)
	go markAsProcessedThenCheckIfCrawlIsDone(url)
}

func checkURL(u url.URL) {
	mux.Lock()
	_, success := checkedURLs[u]
	if !success {
		checkedURLs[u] = false
		go processURL(u)
	}
	mux.Unlock()
}

func listenOnURLCheckingChannelLimited(maxRequests int) {
	requests := 0
	for elem := range urlCheckingChannel {
		requests++
		if requests < maxRequests {
			go checkURL(elem)
		}
	}
}

func listenOnURLCheckingChannel() {
	for elem := range urlCheckingChannel {
		go checkURL(elem)
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
