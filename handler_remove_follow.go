package main

import (
	"fmt"
	"context"
)

func handlerRemoveFollow(s *state, cmd command) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("usage: %s <url>", cmd.command)
	}

	feed, err := s.db.GetFeedsByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error retrieving feed by URL: %w", err)
	}

	err = s.db.RemoveFollow(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}
	
	return nil
}
