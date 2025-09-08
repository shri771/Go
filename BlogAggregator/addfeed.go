package main

import (
	"context"
	"fmt"

	"github.com/shri771/Go/BlogAggregator/internal/rssapi"
)

func handleragg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := rssapi.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("could not list Feed %w", err)
	}

	fmt.Println(feed)
	return nil

}
