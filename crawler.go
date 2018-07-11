package main

import (
	"net/url"
	"sync"
)

func (c *crawler) listenOnURLCheckingChannel() {
	for elem := range c.urlCheckingChannel {
		go c.checkURL(elem)
	}
}

func initCrawler() crawler {
	c := crawler{make(map[url.URL]bool), make(chan url.URL), make(chan page), sync.Mutex{}}
	go c.listenOnURLCheckingChannel()
	return c
}

type crawler struct {
	checkedURLs         map[url.URL]bool
	urlCheckingChannel  chan url.URL
	finishedPageChannel chan page
	mux                 sync.Mutex
}

func (c *crawler) checkURL(u url.URL) {
	c.mux.Lock()
	_, success := c.checkedURLs[u]
	if !success {
		c.checkedURLs[u] = false
		go c.processURL(u)
	}
	c.mux.Unlock()
}

func (c *crawler) processURL(url url.URL) {
	page := parseURL(url)
	for _, link := range page.links {
		c.urlCheckingChannel <- link
	}
	c.finishedPageChannel <- page
	// time.Sleep(time.Millisecond * 20)
	go c.markAsProcessedAndCheckIfDone(url)
}

func (c *crawler) markAsProcessedAndCheckIfDone(url url.URL) {
	c.mux.Lock()
	c.checkedURLs[url] = true
	if c.haveAllQueuedURLsBeenProcessed() {
		close(c.finishedPageChannel)
	}
	c.mux.Unlock()
}

func (c *crawler) haveAllQueuedURLsBeenProcessed() bool {
	for _, x := range c.checkedURLs {
		if !x {
			return false
		}
	}
	return true
}
