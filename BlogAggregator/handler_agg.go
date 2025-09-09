package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shri771/Go/BlogAggregator/internal/rssapi"
)

func handleragg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Enter atlest one argument")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Could not context input to time.Duratino %w", err)
	}

	fmt.Printf("Collecting feeds every %v", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Colud not get next Feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)

	// fmt.Println(addedData)
	feed, err := rssapi.FetchFeed(context.Background(), nextFeed.Url)

	realFeed := feed.Channel

	for _, item := range realFeed.Item {
		fmt.Printf("* Title: %v \n", item.Title)
		fmt.Printf("* Link: %v \n", item.Link)
		fmt.Printf("* Description: %v \n", item.Description)
		fmt.Printf("* PubDate: %v \n", item.PubDate)
	}
	return nil
}
