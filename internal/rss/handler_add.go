package rss

import (
	"context"
	"fmt"
	"os"
	"time"

	cfg "github.com/HemahWeb/gator/internal/config"
	"github.com/HemahWeb/gator/internal/database"
)

func HandlerAdd(state *cfg.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		fmt.Println("Usage: gator rss add <name> <url>")
		os.Exit(1)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := state.DB.AddFeed(context.Background(), database.AddFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Printf("Failed to add feed: %v\n", err)
		os.Exit(1)
	}

	_, err = state.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		fmt.Printf("Failed to create feed follow: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Feed added:")
	fmt.Printf("Feed Name: %s\n", feed.Name)
	fmt.Printf("Feed URL: %s\n", feed.Url)
	fmt.Printf("Feed ID: %d\n", feed.ID)
	fmt.Printf("Feed User ID: %s\n", feed.UserID)
	fmt.Printf("Feed Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Feed Updated At: %s\n", feed.UpdatedAt)

	return nil
}
