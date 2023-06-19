package resolver

import "helper_openai_bot/internal/core/command/infrastructure"

type CommandResolver interface {
	ResolveByCommand(rawCommand string) (*infrastructure.Command, error)
}

type commandResolver struct {
}

func CreateCommandResolver() CommandResolver {
	return &commandResolver{}
}

func (r *commandResolver) ResolveByCommand(rawCommand string) (*infrastructure.Command, error) {
	command, err := infrastructure.TryFrom(rawCommand)
	if err != nil {
		return nil, err
	}

	return command, nil
}
