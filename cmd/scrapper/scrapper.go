package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"golang.org/x/text/encoding/charmap"
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

	songCollector := c.Clone()

	// On every a element which has href attribute call callback
	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// fix regexp or fuck this
		if e.Text == "" {
			return
		}

		// Print link
		decodedSongTitle := decodeWindows1251([]byte(e.Text))
		log.Printf("Song title found: %q\n", decodedSongTitle)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		songCollector.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// On every a element which has `div[id=cont]` attribute call callback
	songCollector.OnHTML(`div[id=cont]`, func(e *colly.HTMLElement) {
		log.Println("Song link found", e.Request.URL)

		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			decodedSmth := decodeWindows1251([]byte(elem.Text))
			log.Printf("Find smth from loop %s", decodedSmth)
		})

		// Find body with lyrics
		dirtyHTML, _ := e.DOM.Html()

		// fixme
		rl := regexp.MustCompile("(</script>)(.+)(<p>)")
		lyricHTML := rl.FindString(dirtyHTML)

		// fixme
		ril := regexp.MustCompile("(</strong>.+</p>)(.+)(<p>)")
		improvedLyricsHTML := ril.FindString(lyricHTML)

		// fixme
		rs := regexp.MustCompile("&nbsp;")
		trimmedHTML := rs.ReplaceAllString(improvedLyricsHTML, "")

		// fixme
		rbr := regexp.MustCompile("<br/><br/>")
		quartedHTML := rbr.Split(trimmedHTML, -1)

		// decodedHTML := decodeWindows1251([]byte(trimmedHTML))
		log.Printf("Find DOM1 %s", quartedHTML[0])
	})

	// fixme
	c.Visit("http://www.gr-oborona.ru/texts/")
}

func decodeWindows1251(ba []uint8) []uint8 {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return out
}
