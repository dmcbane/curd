#! /usr/bin/env bash

function curr() {
  D=$(curd "$@")
  cd "${D}"
}

# Tab completion for curr: suggest saved keywords from `curd ls -k`.
# `complete` is a bash builtin, so only register it when running under bash.
if [ -n "$BASH_VERSION" ]; then
  _curr_complete() {
    COMPREPLY=()
    if [ "$COMP_CWORD" -eq 1 ]; then
      local cur=${COMP_WORDS[COMP_CWORD]}
      COMPREPLY=($(compgen -W "$(curd ls -k)" -- "$cur"))
    fi
    return 0
  }
  complete -F _curr_complete curr
fi
