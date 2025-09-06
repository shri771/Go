package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/BlogAggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("could not create user %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %v switched successfully! \n", name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Delusers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users table: %w", err)
	}

	fmt.Printf("Successfully Reseted Users Table")
	return nil
}

func handlerList(s *state, cmd command) error {
	regUSers, err := s.db.Getusers(context.Background())
	if err != nil {
		return fmt.Errorf("could not list the users: %w", err)
	}

	for _, user := range regUSers {
		fmt.Printf("%v", user)
		if user == s.cfg.CurrentUSerName {
			fmt.Printf(" (current) \n")
		} else {
			fmt.Printf("\n")
		}
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
