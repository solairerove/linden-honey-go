package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/text/encoding/charmap"
)

// Song ... tbd
type Song struct {
	Title  string
	Link   string
	Author string
	Album  string
	Verses []string
}

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

	var song Song

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

		song.Title = string(decodedSongTitle)

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
		// log.Println("Song link found", e.Request.URL)

		song.Link = e.Request.URL.String()

		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			decodedSmth := decodeWindows1251([]byte(elem.Text))
			log.Printf("Find smth from loop %s", decodedSmth)

			// fixme
			if strings.Contains(string(decodedSmth), "Автор") {
				song.Author = string(decodedSmth)
			}

			if strings.Contains(string(decodedSmth), "Альбом") {
				song.Album = string(decodedSmth)
			}
		})

		// Find body with lyrics
		dirtyHTML, _ := e.DOM.Html()

		// fixme
		rl := regexp.MustCompile("(</script>)(.+)(<p>)")
		lyricHTML := rl.FindString(dirtyHTML)

		// fixme
		ril := regexp.MustCompile(`<\/p><p><strong>.+<\/strong>.+<\/p>(?P<Lyrics>.+)<p>`)
		improvedLyricsHTML := ril.FindAllStringSubmatch(lyricHTML, -1)
		names := ril.SubexpNames()

		// if non match patter return
		if improvedLyricsHTML == nil {
			return
		}

		// create map with group name -> content
		md := map[string]string{}
		for i, n := range improvedLyricsHTML[0] {
			md[names[i]] = n
		}

		// split to verses group
		rlp := regexp.MustCompile(`<br/><br/>`)
		unparsedLyrics := rlp.Split(md["Lyrics"], -1)

		// split to separated verses
		verses := make([]string, 0)
		for _, e := range unparsedLyrics {
			log.Print("\n")
			str := regexp.MustCompile(`<br/>`).Split(e, -1)
			for _, s := range str {
				result := regexp.MustCompile(`&#39;`).ReplaceAllString(s, "'")

				// non suka breaking space replaced by human readble space
				trimmedResult := regexp.MustCompile(" ").ReplaceAllString(result, " ")
				decodedResult := decodeWindows1251([]byte(trimmedResult))
				verses = append(verses, string(decodedResult)+"\n")

				log.Printf("Lyrics found %s", string(decodedResult))
			}

			verses = append(verses, "\n\n")
		}

		song.Verses = verses

		log.Printf("Prepare to save next Song -> %s", song)
	})

	// fixme
	c.Visit("http://www.gr-oborona.ru/texts/")
}

// decode shitty cp1251 to human readalbe utf-8
func decodeWindows1251(ba []uint8) []uint8 {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return out
}
