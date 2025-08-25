package rss

import (
	"context"
	"fmt"
	"os"

	cfg "github.com/HemahWeb/gator/internal/config"
	database "github.com/HemahWeb/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *cfg.State, cmd Command, user database.User) error) func(*cfg.State, Command) error {
	return func(s *cfg.State, cmd Command) error {
		user, err := s.DB.GetUserByName(context.Background(), cfg.GetUser())
		if err != nil {
			fmt.Println("You must be logged in to run this command.")
			os.Exit(1)
		}
		return handler(s, cmd, user)
	}
}
