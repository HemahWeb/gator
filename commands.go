package main

import (
	"fmt"

	cfg "github.com/HemahWeb/gator/internal/config"
	"github.com/HemahWeb/gator/internal/rss"
)

type Command = rss.Command

type Commands struct {
	commands map[string]func(state *cfg.State, cmd Command) error
}

func NewCommands() *Commands {
	return &Commands{
		commands: make(map[string]func(state *cfg.State, cmd Command) error),
	}
}

func (c *Commands) run(state *cfg.State, cmd Command) error {
	handler, exists := c.commands[cmd.Name]
	if !exists {
		return fmt.Errorf("command %s not found", cmd.Name)
	}
	return handler(state, cmd)
}

func (c *Commands) register(name string, handler func(state *cfg.State, cmd Command) error) {
	c.commands[name] = handler
}
