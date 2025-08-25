package rss

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
)

func HandlerFeeds(state *cfg.State, cmd Command) error {
	feeds, err := state.DB.ListFeeds(context.Background())
	if err != nil {
		fmt.Printf("Failed to list feeds: %v\n", err)
		os.Exit(1)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.FeedName)
		fmt.Printf("Feed URL: %s\n", feed.FeedUrl)
		fmt.Printf("Created by: %s\n", feed.UserName)
	}
	return nil
}
