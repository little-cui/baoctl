package cmd

import (
	"baoctl/pkg/cmd/command"
	"context"
	"fmt"
)

var manager = New()

type Manager struct {
	cmds []*command.Command
}

func (m *Manager) Add(cmd *command.Command) {
	m.cmds = append(m.cmds, cmd)
}

func (m *Manager) Commands() []*command.Command {
	return m.cmds
}

func (m *Manager) PrintCommands() {
	for _, cmd := range m.cmds {
		fmt.Println(cmd)
	}
}

func (m *Manager) Exec(ctx context.Context, code int, args ...interface{}) error {
	for _, cmd := range m.cmds {
		if cmd.Code == code {
			return cmd.Action(ctx, args...)
		}
	}
	return fmt.Errorf("unexcepted command code")
}

func New() *Manager {
	return &Manager{
		cmds: make([]*command.Command, 0),
	}
}

func Instance() *Manager {
	return manager
}

func RegisterCommand(cmd *command.Command) {
	Instance().Add(cmd)
}
