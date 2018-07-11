package main

import (
	"net/url"
	"testing"
)

func TestHaveAllQueuedURLsBeenProcessed_mix(t *testing.T) {
	// init
	c := initCrawler()

	// arange
	u1, _ := url.Parse("https://monzo.com")
	c.checkedURLs[*u1] = true

	u2, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u2] = true

	u3, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u3] = false

	// act
	output := c.haveAllQueuedURLsBeenProcessed()

	// assert
	testpassed := output == false

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}

func TestHaveAllQueuedURLsBeenProcessed_allTrue(t *testing.T) {
	// init
	c := initCrawler()

	// arange
	u1, _ := url.Parse("https://monzo.com")
	c.checkedURLs[*u1] = true

	u2, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u2] = true

	u3, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u3] = true

	// act
	output := c.haveAllQueuedURLsBeenProcessed()

	// assert
	testpassed := output == true

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}

func TestHaveAllQueuedURLsBeenProcessed_allFalse(t *testing.T) {
	// init
	c := initCrawler()

	// arange
	u1, _ := url.Parse("https://monzo.com")
	c.checkedURLs[*u1] = false

	u2, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u2] = false

	u3, _ := url.Parse("https://monzo.com/blog")
	c.checkedURLs[*u3] = false

	// act
	output := c.haveAllQueuedURLsBeenProcessed()

	// assert
	testpassed := output == false

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}
