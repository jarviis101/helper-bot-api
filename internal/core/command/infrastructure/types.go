package infrastructure

import (
	"fmt"
	"strings"
)

type Command string

const (
	StartCommand  Command = "/start"
	DonateCommand Command = "/donate"
)

var commands = []Command{
	StartCommand,
	DonateCommand,
}

func (c Command) String() string {
	return string(c)
}

func TryFrom(rawCommand string) (*Command, error) {
	for i, command := range commands {
		if strings.EqualFold(command.String(), rawCommand) {
			return &commands[i], nil
		}
	}

	return nil, fmt.Errorf("Unknown command: %s", rawCommand)
}
