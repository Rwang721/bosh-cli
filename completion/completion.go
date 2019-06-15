package completion

type Completion struct {
	Commands    []string
	RootArgs    []string
	CommandArgs map[string][]string
}

func NewCompletion() *Completion {
	return &Completion{
		CommandArgs: make(map[string][]string),
	}
}

func (c *Completion) AddCommand(cmd string) {
	c.Commands = append(c.Commands, cmd)
}

func (c *Completion) AddRootArgs(arg string) {
	c.RootArgs = append(c.RootArgs, arg)
}

func (c *Completion) AddCommandArgs(cmd, arg string) {
	if c.CommandArgs[cmd] == nil {
		c.CommandArgs[cmd] = make([]string, 1)
	}
	c.CommandArgs[cmd] = append(c.CommandArgs[cmd], arg)
}
