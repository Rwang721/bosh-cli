_bosh_complete()
{
  local cur_word prev_word subcmd type_list

  # COMP_WORDS is an array of words in the current command line.
  # COMP_CWORD is the index of the current word (the one the cursor is
  # in). So COMP_WORDS[COMP_CWORD] is the current word; we also record
  # the previous word here, although this specific script doesn't
  # use it yet.
  cur_word="${COMP_WORDS[COMP_CWORD]}"
  prev_word="${COMP_WORDS[COMP_CWORD-1]}"

  # Only perform completion if the current word starts with a dash ('-') or
  # blank, and the previous word is "bosh". This means that that the user is
  # trying to complete an option or command, respectively.
  #
  # Otherwise, treat the second word in the current command line as a
  # subcommand to `bosh`.
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
        type_list="$(bosh completion -c ${subcmd})"
      fi
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
complete -d -F _bosh_complete bosh
