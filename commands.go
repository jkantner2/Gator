package main

import (
	"fmt"
	"strings"
)

type command struct {
	command		string
	arguments	[]string
}

type commands struct {
	registeredCommands	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if c.registeredCommands[cmd.command] == nil {
		return fmt.Errorf("Invalid command entered\n")
	}
	function := c.registeredCommands[cmd.command]
	return function(s, cmd)
}

//register function
//clean name string... no leading/training whitespace
//add name to 
func (c *commands) register(name string, f func(*state, command)error) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("No command name entered\n")
	}
	c.registeredCommands[name] = f
	return nil
}
