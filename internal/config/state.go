package config

import (
	database "github.com/HemahWeb/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *Config
}

func NewState() *State {
	cfg := GetConfig()
	return &State{
		Config: &cfg,
	}
}
