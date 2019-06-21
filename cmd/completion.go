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
	switch {
	case opts.RootArgs:
		c.printBoshOpts()
	case opts.CommandArgs != "":
		c.printArgsForSubcommand(opts.CommandArgs)
	case opts.Commands:
		c.printBoshCommands()
	case opts.Export:
		fallthrough
	default:
		fmt.Print(completionScript())
	}

	return nil
}

func (c CompletionCmd) printBoshOpts() {
	for _, v := range c.compl.RootArgs {
		fmt.Println(v)
	}
}

func (c CompletionCmd) printArgsForSubcommand(cmd string) {
	for _, v := range c.compl.CommandArgs[cmd] {
		fmt.Println(v)
	}
}

func (c CompletionCmd) printBoshCommands() {
	for _, v := range c.compl.Commands {
		fmt.Println(v)
	}
}

func completionScript() string {
	return `
_bosh_complete()
{
  local cur_word prev_word subcmd type_list

  # COMP_WORDS is an array of words in the current command line.
  # COMP_CWORD is the index of the current word (the one the cursor is
  # in). So COMP_WORDS[COMP_CWORD] is the current word.
  cur_word="${COMP_WORDS[COMP_CWORD]}"
  prev_word="${COMP_WORDS[COMP_CWORD-1]}"

  # Only perform completion if the current word starts with a dash ('-') or
  # blank, and the previous word is "bosh". This means that that the user is
  # trying to complete an option or command, respectively.
  #
  # Otherwise, treat the second word in the current command line as a
  # subcommand to 'bosh'.
  for word in "${COMP_WORDS[@]}"; do
    if [[ "${word}" != "bosh" && ! "${word}" =~ ^\- ]]; then
      subcmd="${word}"
    fi
  done

  case "${cur_word}" in
    -*)
      if [[ -z "${subcmd}" ]]; then
        type_list="$(bosh completion -r)"
      elif [[ -n "${subcmd}" ]]; then
        type_list="$(bosh completion --command ${subcmd})"
      fi
      ;;
    *)
      type_list="$(bosh completion --commands)"
      ;;
  esac

  local OLDIFS="$IFS"
  local IFS=$'\n'

  # COMPREPLY is the array of possible completions, generated with
  # the compgen builtin.
  COMPREPLY=( $(compgen -W "${type_list}" -- "${cur_word}") )

  IFS="$OLDIFS"
  if [[ ${#COMPREPLY[*]} -eq 1 ]]; then #Only one completion
    COMPREPLY=( ${COMPREPLY[0]%% - *} ) #Remove ' - ' and everything after
  fi

  return 0
}

# Register _bosh_complete to provide completion for the following commands
complete -d -F _bosh_complete bosh
`
}
