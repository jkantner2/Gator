package main

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"
	"github.com/jkantner2/Gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {	
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.command)
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]


	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
		Url:		url,
		UserID:		user.ID,
	})
	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID:			feed.ID,
		CreatedAt:	feed.CreatedAt,
		UpdatedAt:	feed.UpdatedAt,
		UserID:		feed.UserID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Couldn't follow feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("===========================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* Created at:	%s\n", feed.CreatedAt)
	fmt.Printf("* Updated at:	%s\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
}
