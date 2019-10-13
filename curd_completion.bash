#/usr/bin/env bash

## function echo-err()
## {
##   echo "$@" >&2
## }

_completions_curd()
{
  COMPREPLY=()   # Array variable storing the possible completions.

  # keep the suggestions in a local variable
  local cur=${COMP_WORDS[COMP_CWORD]}
  local suggestions=($(compgen -W "$(curd -C ${COMP_WORDS[@]})" -- "$cur"))
  COMPREPLY=("${suggestions[@]}")

  return 0
}

complete -F _completions_curd curd

_completions_curr()
{
  COMPREPLY=()   # Array variable storing the possible completions.

  # keep the suggestions in a local variable
  local suggestions

  if [ ${#COMP_WORDS[@]} -le 2 ]; then
    suggestions=($(compgen -W "$(curd ls -k)" -- "${COMP_WORDS[1]}"))
  else
    return 0
  fi
  COMPREPLY=("${suggestions[@]}")

  return 0
}

complete -F _completions_curr curr
