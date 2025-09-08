package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/BlogAggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("Entere  two arguments")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get Current users id
	ctx := context.Background()

	// Add Feed
	feed, err := s.db.AddFeed(ctx, database.AddFeedParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       url,
	})
	if err != nil {
		// return fmt.Errorf("could not Add feed to database: %w", err)

	}

	// Create feed follow

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeed(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.ID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}

func getCurrentUserId(s *state) (uuid.UUID, error) {
	currentUser := s.cfg.CurrentUSerName

	id, err := s.db.GetUserid(context.Background(), currentUser)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Could not Get the currentUser id: %w", err)
	}
	return id, nil
}
