_bosh_complete()
{
  local cur_word prev_word type_list

  # COMP_WORDS is an array of words in the current command line.
  # COMP_CWORD is the index of the current word (the one the cursor is
  # in). So COMP_WORDS[COMP_CWORD] is the current word; we also record
  # the previous word here, although this specific script doesn't
  # use it yet.
  cur_word="${COMP_WORDS[COMP_CWORD]}"
  prev_word="${COMP_WORDS[COMP_CWORD-1]}"

  # Only perform completion if the current word starts with a dash ('-'),
  # meaning that the user is trying to complete an option.
  case "${cur_word}" in
    -*)
      # Generate a list of types it supports
      type_list="$(bosh completion -r)"
      ;;
    *)
      type_list="$(bosh completion)"
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
complete -F _bosh_complete bosh
