package main

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"
	"Gator/internal/database"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.command)
	}
	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal()
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.command)
	}

	name := cmd.arguments[0]

	userParams := database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		name,
	}

	fmt.Println(userParams.Name)

	_, err := s.db.GetUser(context.Background(), userParams.Name)
	if err == nil {
		log.Fatal()
		return err
	}

	_, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	s.cfg.SetUser(userParams.Name)
	fmt.Printf("User: %s was created\n", userParams.Name)
	fmt.Printf("user data: %v", userParams)

	return nil
}
func handlerUsers(s *state, cmd command) error {
	users := []database.User{}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error retreiving users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.Current_username{
			fmt.Printf("%s (current)\n", user.Name)
			continue
		}

		fmt.Println(user.Name)
	}

	return nil
}

