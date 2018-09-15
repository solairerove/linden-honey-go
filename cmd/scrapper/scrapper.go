package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domain: www.gr-oborona.ru
		colly.AllowedDomains("www.gr-oborona.ru"),

		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),

		// Visit only root url and urls which start with "text" on www.gr-oborona.ru
		colly.URLFilters(
			regexp.MustCompile("http://www.gr-oborona.ru/texts/"), // fixme
		),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://www.gr-oborona.ru/texts/")
}
