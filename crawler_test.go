package main

import (
	"net/url"
	"testing"
)

func TestHaveAllQueuedURLsBeenProcessed_mix(t *testing.T) {
	// init
	checkedURLs := make(map[url.URL]bool)

	// arange
	u1, _ := url.Parse("https://www.totaljobs.com")
	checkedURLs[*u1] = true

	u2, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u2] = true

	u3, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u3] = false

	// act
	output := haveAllQueuedURLsBeenProcessed(checkedURLs)

	// assert
	testpassed := output == false

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}

func TestHaveAllQueuedURLsBeenProcessed_allTrue(t *testing.T) {
	// init
	checkedURLs := make(map[url.URL]bool)

	// arange
	u1, _ := url.Parse("https://www.totaljobs.com")
	checkedURLs[*u1] = true

	u2, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u2] = true

	u3, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u3] = true

	// act
	output := haveAllQueuedURLsBeenProcessed(checkedURLs)

	// assert
	testpassed := output == true

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}

func TestHaveAllQueuedURLsBeenProcessed_allFalse(t *testing.T) {
	// init
	checkedURLs := make(map[url.URL]bool)

	// arange
	u1, _ := url.Parse("https://www.totaljobs.com")
	checkedURLs[*u1] = false

	u2, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u2] = false

	u3, _ := url.Parse("https://www.totaljobs.com/blog")
	checkedURLs[*u3] = false

	// act
	output := haveAllQueuedURLsBeenProcessed(checkedURLs)

	// assert
	testpassed := output == false

	if !testpassed {
		t.Error("haveAllQueuedURLsBeenProcessed() produced an unexpected result")
	}
}
