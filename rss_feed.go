package main

import (
	"fmt"
	"context"
	"io"
	"net/http"
	"encoding/xml"
	"html"
	"time"
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

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 5*time.Second,
	}
	
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error generating HTTP request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error generating response: %w", err)
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error reading response: %w", err)
	}

	RSSFeedResp := RSSFeed{}
	err = xml.Unmarshal(dat, &RSSFeedResp)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error Unmarshalling xml: %w", err)
	}

	RSSFeedResp.Channel.Title = html.UnescapeString(RSSFeedResp.Channel.Title)
	RSSFeedResp.Channel.Description = html.UnescapeString(RSSFeedResp.Channel.Description)

	for i, item := range RSSFeedResp.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		RSSFeedResp.Channel.Item[i] = item
	}

	return &RSSFeedResp, nil
}
