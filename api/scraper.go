package api

import (
	"log"
	"time"

	"github.com/Jhnvlglmlbrt/rss-aggregator/internal/database"
)

func startScraping(
	db *database.Queries, // connection to database
	concurrency int, // number of diff goroutines to do the scraping
	timeBetweenRequest time.Duration, // time between each request to go scrape a new rss feed
) {
	log.Printf("Scraping on %v goroutines every %s duraiton", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {

	}
}
