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
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if token.DataAtom == atom.A {
			for i := range token.Attr {
				if token.Attr[i].Key == "href" {
					href := token.Attr[i].Val
					if len(href) > 1 && href[0] == "/"[0] {
						x, err := url.Parse(href)
						x.Host = page.url.Host
						x.Fragment = ""
						if err == nil {
							page.links = append(page.links, *x)
						}
						// fmt.Println(token.Attr[i].Val)
						// fmt.Println(token.String())
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
