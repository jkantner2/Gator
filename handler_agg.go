package main

import (
	"fmt"
	"context"
	"github.com/jkantner2/Gator/internal/database"
	"github.com/google/uuid"
	"log"
	"time"
	"database/sql"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) < 1 || len(cmd.arguments) > 2 {
		return fmt.Errorf("usage: %v <time_between_requests>", cmd.command)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Error parsing time between requets ensure proper usage: %s <timeBetweenRequests> ie 1s, 1m, 1h. \n %w", cmd.command, err)
	}

	log.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Error fetching next feed: %w", err)
		return
	}
	log.Println("Found a feed to fetch")
	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v\n", feed.Name, err)
		return
	}

	feedData, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Error fetching rss feed %s: %v\n", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item{
		layouts := []string{
			time.RFC822,
			time.RFC1123,
			time.RFC1123Z,
			time.RFC822Z,
		}

		var parsedTime time.Time
		var err error
		for _, layout := range layouts {
			parsedTime, err = time.Parse(layout, item.PubDate)
			if err == nil {
				break
			}
		}

		db.CreatePost(context.Background(), database.CreatePostParams{
			ID:			 uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{DerefOrEmpty(&item.Title), IsNotNil(&item.Title)},
			Url:         item.Link,
			Description: sql.NullString{DerefOrEmpty(&item.Description), IsNotNil(&item.Description)},
			PublishedAt: sql.NullTime{parsedTime, err == nil},
			FeedID:      feed.ID,
		})
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))	
}

func DerefOrEmpty[T any](val *T) T {
	if val == nil { 
	var empty T
		return empty
	}
	
	return *val
}

func IsNotNil[T any](val *T) bool {
	return val != nil
}
