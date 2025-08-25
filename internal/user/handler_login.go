package user

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	"github.com/HemahWeb/gator/internal/rss"
)

func HandlerLogin(state *cfg.State, cmd rss.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username provided")
	}
	username := cmd.Args[0]

	if _, err := state.DB.GetUserByName(context.Background(), username); err != nil {
		fmt.Printf("User %s not found\n", username)
		os.Exit(1)
	}

	cfg.SetUser(username)

	config := cfg.GetConfig()
	state.Config = &config

	fmt.Printf("Logged in as %s\n", username)

	return nil
}
