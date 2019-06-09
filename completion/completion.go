package completion

type Completion struct {
	Commands    []string
	RootArgs    []string
	CommandArgs []string
}

func NewCompletion() *Completion {
	return &Completion{}
}

func (c *Completion) AddCommand(cmd string) {
	c.Commands = append(c.Commands, cmd)
}

func (c *Completion) AddRootArgs(arg string) {
	c.RootArgs = append(c.RootArgs, arg)
}

func (c *Completion) AddCommandArgs(arg string) {
	c.CommandArgs = append(c.CommandArgs, arg)
}
