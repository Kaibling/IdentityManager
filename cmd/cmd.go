package cmd

import "errors"

type Command struct {
	Name string
	Args []string
	F    func([]string) error
}

func (c *Command) Execute(args []string) error {
	return c.F(args)
}

type Commands struct {
	Commands map[string]Command
}

func (c *Commands) AddCommand(name string, f func([]string) error) {
	c.Commands[name] = Command{Name: name, F: f}
}
func (c *Commands) Exec(args []string) error {

	commandName := args[0]
	if cmd, ok := c.Commands[commandName]; ok {
		return cmd.Execute(args[1:])
	}
	return errors.New("no Command found")
}

func NewCommands() *Commands {
	return &Commands{Commands: map[string]Command{}}
}
