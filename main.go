package main

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

const BaseURL = "https://www.nytimes.com/section/business/"

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		isAbsolutePath := strings.HasPrefix(e.Attr("href"),
			"/") && e.Attr("href") != "/"
		hasBaseURL := strings.HasPrefix(e.Attr("href"), BaseURL)

		if isAbsolutePath || hasBaseURL {
			log.Println("Found:", e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	if err := c.Visit(BaseURL); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}