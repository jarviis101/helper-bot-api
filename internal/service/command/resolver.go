package command

import "helper_openai_bot/internal/service/command/infrastructure"

type commandResolver struct {
}

func CreateCommandResolver() infrastructure.CommandResolver {
	return &commandResolver{}
}

func (r *commandResolver) ResolveByCommand(rawCommand string) (*infrastructure.Command, error) {
	command, err := infrastructure.TryFrom(rawCommand)
	if err != nil {
		return nil, err
	}

	return command, nil
}
