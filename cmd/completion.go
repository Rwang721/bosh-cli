package cmd

import (
	"fmt"
)

var commands []string

type CompletionCmd struct {
}

func NewCompletionCmd() CompletionCmd {
	return CompletionCmd{}
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
	fmt.Println("--foo")
	fmt.Println("--bar")
	fmt.Println("--baz")
}

func (c CompletionCmd) printBoshCommands() {
	for _, v := range commands {
		fmt.Println(v)
	}
}

func AddCommandForCompletion(name string) {
	commands = append(commands, name)
}
