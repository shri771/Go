package rssapi

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// New Client
	timeout := 5 * time.Second
	httpClient := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	// Set headers
	req.Header.Add("User-Agent", "gator")

	// Make request
	res, err := httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	// Read data
	buf, err := io.ReadAll(res.Body)

	// Decode the data
	feed := RSSFeed{}
	err = xml.Unmarshal(buf, &feed)
	if err != nil {
		return &RSSFeed{}, err
	}

	return cleanData(&feed), nil

}

func cleanData(feed *RSSFeed) *RSSFeed {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	// clean the RSSItem
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
	return feed
}
