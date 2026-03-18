package main

import (
	"fmt"
	"github.com/jkantner2/Gator/internal/database"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.command)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error querying feeds: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUsername(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error querying user associated with creation of feed %w", err)
		}
		printFeedInfo(feed)
		fmt.Printf("* Username:			%s\n", username)

	}
	return nil
}

func printFeedInfo(feed database.Feed) {
	fmt.Printf("* Feed Name:			%s\n", feed.Name)
	fmt.Printf("* URL:					%s\n", feed.Url)
}
