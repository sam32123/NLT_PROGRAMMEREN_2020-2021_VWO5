package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.ah.nl/allerhande/recept/R-R1194290/makkelijke-risotto")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("itemprop")
		if attr == "ingredients" {
			fmt.Println(strings.Replace(s.Text(), "\n", "", -1))
		}
	})
}

func main() {
	ExampleScrape()
}
