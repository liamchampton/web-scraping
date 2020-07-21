package actions

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
)

func Crawl(w http.ResponseWriter, r *http.Request) {

	// Array containing all the known URLs in a sitemap
	var knownUrls []string

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains("www.shopify.com"))

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Start the collector
	c.Visit("https://www.shopify.com/sitemap.xml")

	fmt.Println("All known URLs:")

	numURLs := strconv.Itoa(len(knownUrls))
	fmt.Println("Collected", numURLs, "URLs")

	knownUrls = append(knownUrls, "Collected "+numURLs+" URLs")

	urlStr := strings.Join(knownUrls, "\n")
	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(urlStr)))
}
