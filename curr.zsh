function curr() {
  local D
  D=$(curd "$@")
  cd "${D}"
}

# Tab completion for curr: suggest saved keywords from `curd ls -k`.
# Requires the zsh completion system; ensure `autoload -U compinit && compinit`
# has run earlier in your ~/.zshrc.
_curr_complete() {
  local -a keywords
  keywords=(${(z)"$(curd ls -k)"})
  compadd -a keywords
}
compdef _curr_complete curr
