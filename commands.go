package main

import (
	"errors"
)

type Commands struct {
	cmds map[string]func(*State, []string) error
}

func (c *Commands) register(name string, f func(*State, []string) error) {
	c.cmds[name] = f
}

func (c *Commands) run(s *State, name string, args []string) error {
	f, found := c.cmds[name]
	if !found {
		return errors.New("command not found")
	}

	return f(s, args)
}
