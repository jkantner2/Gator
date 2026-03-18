package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.command)
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting users table: %w", err)
	}
	
	fmt.Println("Database reset successfully!")
	return nil
}
