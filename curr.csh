# csh/tcsh has no shell functions, so curr is an alias. \!* forwards all
# arguments to curd, and the backquotes capture the resulting path to cd into.
alias curr 'cd "`curd \!*`"'

# Tab completion for curr: suggest saved keywords from `curd ls -k`.
# The `complete` builtin is a tcsh feature (plain csh ignores it).
# p/1/.../ completes the first positional argument from the command output.
complete curr 'p/1/`curd ls -k`/'
