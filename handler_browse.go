package main

import (
	"context"
	"fmt"
	"strconv"
	"github.com/jkantner2/Gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	if len(cmd.arguments) < 1 || len(cmd.arguments) > 2 {
		return fmt.Errorf("usage: %v <max_results>", cmd.command)
	}

	limit, err := strconv.Atoi(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Failed to convert limit string integer: %v", err)
	}
	
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_username)	
	if err != nil {
		return fmt.Errorf("Error obtaining current user info: %v", err)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{user.ID, int32(limit)})
	if err != nil {
		return fmt.Errorf("Error retrieving posts: %v", err)	
	}

	for _, post := range posts {
		fmt.Printf("Publish Date: %v\n Title: %s\n URL: %s\n Description: %s\n", post.PublishedAt, post.Title.String, post.Url, post.Description.String)
	}

	return nil
}
