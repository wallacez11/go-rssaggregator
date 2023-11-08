package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/wallacez11/go-rssaggregator/internal/database"
)

func startScraping(db *database.Queries, concurrrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFatch(context.Background(), int32(concurrrency))
		if err != nil {
			log.Println("error fetching feeds: ", err)
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}

		wg.Wait()
	}

}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
		return
	}

	RSSFeed, err := urlTofeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range RSSFeed.Channel.Item {
		log.Println("Found post", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(RSSFeed.Channel.Item))
}
