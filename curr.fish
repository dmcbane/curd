function curr
    set -l D (curd $argv)
    cd "$D"
end

# Tab completion for curr: suggest saved keywords from `curd ls -k`.
# -f disables file completion so only keywords are offered.
complete -c curr -f -a '(curd ls -k | string split -n " ")'
