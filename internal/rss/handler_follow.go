package rss

import (
	"context"
	"fmt"
	"os"
	"time"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func HandlerFollow(state *cfg.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		fmt.Printf("No feed url provided\n")
		os.Exit(1)
	}

	feedURL := cmd.Args[0]

	feed, err := state.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		fmt.Printf("Feed does not exist: %v\n", err)
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

	fmt.Printf("%s is now following:\n %s\n", user.Name, feed.Name)

	return nil
}
