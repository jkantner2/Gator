package main

import (
	"fmt"
	"context"
	"github.com/jkantner2/Gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.command)
	}

	following, err := s.db.GetFollowing(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Could't retrieve followed feeds: %w", err)
	}

	printFollowing(following)
	return nil
}

func printFollowing(following []database.GetFollowingRow) {
	for i, _ := range following {
		
		fmt.Printf("Feed name: %s\n", following[i].FeedName)
	}
}
