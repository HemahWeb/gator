package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
	rss "github.com/HemahWeb/gator/internal/rss"
	user "github.com/HemahWeb/gator/internal/user"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		fmt.Println("Available commands:")
		fmt.Println("  login <username>")
		fmt.Println("  register <username>")
		fmt.Println("  reset")
		fmt.Println("  users")
		fmt.Println("  addfeed <name> <url>")
		os.Exit(1)
	}

	dbURL := cfg.GetConfig().DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	state := cfg.NewState()
	state.DB = database.New(db)

	commands := NewCommands()
	commands.register("login", user.HandlerLogin)
	commands.register("register", user.HandlerRegister)
	commands.register("reset", user.HandlerReset)
	commands.register("users", user.GetUsers)

	commands.register("agg", rss.Agg)
	commands.register("addfeed", rss.MiddlewareLoggedIn(rss.HandlerAdd))
	commands.register("feeds", rss.HandlerFeeds)
	commands.register("following", rss.MiddlewareLoggedIn(rss.HandlerFollowing))
	commands.register("follow", rss.MiddlewareLoggedIn(rss.HandlerFollow))
	commands.register("unfollow", rss.MiddlewareLoggedIn(rss.HandlerDeleteFeedFollow))
	commands.register("browse", rss.MiddlewareLoggedIn(rss.HandlerBrowse))

	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	err = commands.run(state, rss.Command{Name: os.Args[1], Args: args})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
