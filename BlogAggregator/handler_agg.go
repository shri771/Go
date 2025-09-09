package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/shri771/Go/BlogAggregator/internal/database"
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

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Colud not get next Feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		ID: nextFeed[0].ID,
	})

	feed, err := rssapi.FetchFeed(context.Background(), nextFeed[0].Url)

	realFeed := feed.Channel
	fmt.Printf("*** %v ***", realFeed.Title)

	for _, item := range realFeed.Item {
		fmt.Printf("* Title: %v \n", item.Title)
		fmt.Printf("* Link: %v \n", item.Link)
		fmt.Printf("* Description: %v \n", item.Description)
		fmt.Printf("* PubDate: %v \n", item.PubDate)

	}

	return nil
}
