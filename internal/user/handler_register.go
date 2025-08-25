package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
	"github.com/HemahWeb/gator/internal/rss"
)

func HandlerRegister(state *cfg.State, cmd rss.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username provided")
	}

	username := cmd.Args[0]

	if _, err := state.DB.GetUserByName(context.Background(), username); err == nil {
		fmt.Printf("User %s already exists\n", username)
		os.Exit(1)
	}

	user, err := state.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	cfg.SetUser(user.Name)
	fmt.Printf("Registered as %s\n", user.Name)

	return nil
}
