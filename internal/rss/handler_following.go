package rss

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func HandlerFollowing(state *cfg.State, cmd Command, user database.User) error {

	feedFollows, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("Failed to get feed follows for user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("You are following:\n")
	for _, feedFollow := range feedFollows {
		feed, err := state.DB.GetFeedByID(context.Background(), feedFollow.FeedID)
		if err != nil {
			fmt.Printf("  (error fetching feed info: %v)\n", err)
			continue
		}
		fmt.Printf("%s\n", feed.Name)
	}

	return nil
}
