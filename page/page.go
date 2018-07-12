package page

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ParseURL(url url.URL) Page {
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println("big problem with " + url.String())
		return Page{}
	}
	page := Page{url: url}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		// move the interator forward
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			// if the token is an ErrorToken then we can assume that this is the end of the file and we need to break the loop
			break
		}
		token := tokenizer.Token()
		// this checks if the token is an <a> tag
		// TODO: add checks for img tags etc
		if token.DataAtom == atom.A {
			for i := range token.Attr {
				// we can get the relevant link from the href attribute
				if token.Attr[i].Key == "href" {
					href := token.Attr[i].Val
					// if the href starts with a / then it's an internal link
					if len(href) > 1 && href[0] == "/"[0] {
						link, err := url.Parse(href)
						// set the host to the same host as the page
						link.Host = page.url.Host
						// set the fragment to an empty string
						link.Fragment = ""
						if err == nil {
							page.links = append(page.links, *link)
						}
					}
				}
			}
		}
	}
	return page
}

type Page struct {
	url   url.URL
	links []url.URL
}

func (p *Page) GetLinks() []url.URL {
	return p.links
}

func (p *Page) PrintPage() {
	fmt.Println(p.url.String())
	for _, link := range p.links {
		fmt.Println(" -- " + link.String())
	}
}
