package main

import (
	"fmt"
	"context"
	"github.com/jkantner2/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.Current_username)
		if err != nil {
			return fmt.Errorf("Couldn't retrieve user: %w", err)
		}

		return handler(s, cmd, currentUser)
	}
}
