package user

import (
	"context"
	"fmt"

	cfg "github.com/HemahWeb/gator/internal/config"
	"github.com/HemahWeb/gator/internal/rss"
)

func GetUsers(state *cfg.State, cmd rss.Command) error {
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	fmt.Println("Users:")
	for _, user := range users {
		if user.Name == cfg.GetConfig().CurrentUserName {
			fmt.Println("* " + user.Name + " (current)")
		} else {
			fmt.Println("* " + user.Name)
		}
	}
	return nil
}
