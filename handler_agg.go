package main

import (
	"fmt"
	"context"
	"github.com/jkantner2/Gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		fmt.Errorf("Error fetching rrs feed: %w", err)
	}
	
	fmt.Printf("RSSFeed: %+v\n", feed)
	
	return nil
}

func scrapeFeeds() {
	nextFeed, err := 
}
