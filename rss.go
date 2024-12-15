package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Gator")

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", result.Status)
	}

	// Read the entire response body
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the XML into the RSSFeed struct
	feed := &RSSFeed{}
	if err := xml.Unmarshal(body, feed); err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for index, rssItem := range feed.Channel.Item {
		feed.Channel.Item[index].Title = html.UnescapeString(rssItem.Title)
		feed.Channel.Item[index].Description = html.UnescapeString(rssItem.Description)
	}

	return feed, nil
}
