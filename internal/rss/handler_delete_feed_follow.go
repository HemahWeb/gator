// Add a new unfollow command that accepts a feed's URL as an argument and unfollows it for the current user. This is, of course, a "logged in" command - use the new middleware.

package rss

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func HandlerDeleteFeedFollow(state *cfg.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		fmt.Println("Usage: gator unfollow <feed_url>")
		os.Exit(1)
	}

	feedURL := cmd.Args[0]

	feed, err := state.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		fmt.Printf("Feed does not exist: %v\n", err)
		os.Exit(1)
	}

	err = state.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		fmt.Printf("Failed to delete feed follow: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Unfollowed %s\n", feed.Name)

	return nil
}
