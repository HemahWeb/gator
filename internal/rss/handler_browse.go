package rss

import (
	"context"
	"fmt"
	"os"
	"strconv"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func HandlerBrowse(state *cfg.State, cmd Command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			fmt.Println("invalid limit")
			os.Exit(1)
		}
		limit = int32(parsedLimit)
	}
	posts, err := state.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		fmt.Println("failed to get posts")
		os.Exit(1)
	}
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
