package rss

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func Agg(state *cfg.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("time between requests is required")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every: %s\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for range ticker.C {
		scrapeFeeds(state)
	}
	return nil
}

func scrapeFeeds(state *cfg.State) {
	feed, err := state.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("no feeds to fetch: %v\n", err)
		return
	}
	err = state.DB.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("could not mark %s as fetched: %v\n", feed.Name, err)
		return
	}
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("could not fetch feed: %v\n", err)
		return
	}

	items := fetchedFeed.Channel.Item
	for _, item := range items {
		_, err := state.DB.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: parsePubDate(item.PubDate), Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			} else {
				fmt.Printf("could not create post %s: %v\n", item.Title, err)
			}
		}
	}
}

func parsePubDate(pubDate string) time.Time {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, pubDate); err == nil {
			return t
		}
	}

	return time.Time{}
}
