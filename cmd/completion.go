package cmd

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/completion"
)

type CompletionCmd struct {
	compl *completion.Completion
}

func NewCompletionCmd(c *completion.Completion) CompletionCmd {
	return CompletionCmd{compl: c}
}

func (c CompletionCmd) Run(opts CompletionOpts) error {
	if opts.Root {
		c.printBoshOpts()
	} else {
		c.printBoshCommands()
	}

	return nil
}

func (c CompletionCmd) printBoshOpts() {
	for _, v := range c.compl.RootArgs {
		fmt.Println(v)
	}
}

func (c CompletionCmd) printBoshCommands() {
	for _, v := range c.compl.Commands {
		fmt.Println(v)
	}
}
