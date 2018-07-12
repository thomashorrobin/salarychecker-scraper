package page

import (
	"net/url"
	"testing"
)

func TestParseUrl(t *testing.T) {
	u1, _ := url.Parse("https://www.totaljobs.com")
	page := ParseURL(*u1)
	success := false
	for okay := range page.links {
		t.Log(okay)
		success = true
	}
	if !success {
		t.Error("no links have been added to a page known to have many links")
	}
}
