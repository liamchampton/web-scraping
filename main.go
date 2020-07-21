package main

import (
	"fmt"
	logr "github.com/sirupsen/logrus"
	"github.com/web-scraping/pkg/actions"

	"net/http"
)

func main() {

	// Create the first route handler listening on '/'
	http.HandleFunc("/", handler)

	http.HandleFunc("/scrape", actions.Scrape)
	http.HandleFunc("/crawl", actions.Crawl)

	logr.Info("Starting up on 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, welcome to your app. Use the folling suffix's on the URL to show the different results.\n1)'/scrape' to show results of web scraping.\n2)'/crawl' to show results of a web crawler"

	// Logs a message to the terminal using the 3rd party import logrus
	logr.Info("Received request for the home page")

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(msg)))
}