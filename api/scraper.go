package api

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Jhnvlglmlbrt/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func StartScraping(
	db *database.Queries, // connection to database
	concurrency int, // number of diff goroutines to do the scraping
	timeBetweenRequest time.Duration, // time between each request to go scrape a new rss feed
) {
	log.Printf("Scraping on %v goroutines every %s duraiton", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency), // limit
		)
		if err != nil {
			log.Println("error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched")
		return
	}

	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("error parsing date %v with err %v", item.PubDate, err)
			continue
		}

		// log.Println("Found post", item.Title, "on feed ", feed.Name)
		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			if strings.Contains(err.Error(), "повторяющееся значение ключа") {
				continue
			}
			log.Println("error creating post: ", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
