package rss

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"

	"github.com/HemahWeb/gator/internal/types"
)

type RSSFeed = types.RSSFeed
type RSSItem = types.RSSItem

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, errors.New("failed to create request, maybe the URL is invalid")
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to get response")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	var feed RSSFeed
	err = xml.Unmarshal(bodyBytes, &feed)
	if err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &feed, nil
}
