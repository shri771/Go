package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/BlogAggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Entere atleast two argumets")
	}

	url := cmd.Args[0]

	// Find user ID
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUSerName)
	if err != nil {
		return fmt.Errorf("Could not retriew the user: %w", err)
	}

	// Find feed id
	feed_id, err := s.db.GetFeedIdByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get the feed_id: %w", err)
	}

	feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed_id,
	})

	fmt.Printf("Successfully Added Feed \n")
	fmt.Printf("* Current User: %v \n", feed.UserName)
	fmt.Printf("* Created Feed is: %v \n", feed.FeedName)

	return nil
}

func handlerFollwing(s *state, cmd command) error {
	follwing, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUSerName)
	if err != nil {
		return fmt.Errorf("could not Get the feed: %v \n", err)
	}

	// fmt.Printf("Follwing For user %v: \n", s.cfg.CurrentUSerName)
	for _, feed := range follwing {
		fmt.Println(feed)
	}
	return nil

}
