package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	remoteFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, rssItem := range remoteFeed.Channel.Item {
		fmt.Println(rssItem.Title)
	}

	return nil
}
