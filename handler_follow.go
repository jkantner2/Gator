package main

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"github.com/jkantner2/Gator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.command)
	}

	feed, err := s.db.GetFeedsByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error retrieving feed by URL: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Feed Name: %s, Current User: %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}
