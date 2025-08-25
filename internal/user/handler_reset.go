package user

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	"github.com/HemahWeb/gator/internal/rss"
)

func HandlerReset(state *cfg.State, cmd rss.Command) error {
	err := state.DB.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("failed to reset users: %v", err)
		os.Exit(1)
	}
	fmt.Println("Users reset")
	os.Exit(0)
	return nil
}
