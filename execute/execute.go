package execute

import (
	"errors"
	"fmt"
	"os"
	"sort"
        "strings"

	"github.com/dmcbane/curd/args"
	"github.com/dmcbane/curd/config"
)

func ExecuteCommand(a args.Args, c config.Config) error {
	switch {
	case a.Clean:
		{
			for k, v := range c.Paths {
				if _, err := os.Stat(v); err != nil && os.IsNotExist(err) {
					delete(c.Paths, k)
				}
			}
			if err := c.WriteConfig(); err != nil {
				return err
			}
		}
        case a.Completion:
                {
                        GenerateBashCompletionWordLists(a.Cmdline)
                }
	case a.List:
		{
			// sort the keys of the arguments map
			var keys []string
			for k, _ := range c.Paths {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			if a.KeywordsOnly {
                                // TODO: skip <default>""
                                result := strings.Join(keys[1:], "  ")
                                fmt.Println(result)
			} else {
			        for _, v := range keys {
					fmt.Printf("%s - %s\n", v, (c.Paths)[v])
				}
			}
		}

	case a.Remove:
		{
			delete(c.Paths, a.Keyword)
			if err := c.WriteConfig(); err != nil {
				return err
			}
		}
	case a.Save:
		{
			if a.Directory == "" {
				pwd, err := os.Getwd()
				if err != nil {
					return err
				}
				(c.Paths)[a.Keyword] = pwd
			} else {
				(c.Paths)[a.Keyword] = a.Directory
			}
			if err := c.WriteConfig(); err != nil {
				return err
			}
		}
	default: // a.Read
		{
			if val, exists := (c.Paths)[a.Keyword]; exists {
				fmt.Println(val)
			} else {
				return errors.New(fmt.Sprintf("%s does not exist in %s", a.Keyword, c.ConfigFile))
			}
		}
	}
	return nil
}

func GenerateBashCompletionWordLists(cmdline []string) {
  for _, s := range cmdline {
    fmt.Printf("%s\n", s)
  }


  //   COMPREPLY=()   # Array variable storing the possible completions.
  //
  //   # keep the suggestions in a local variable
  //   local suggestions
  //   local cur=${COMP_WORDS[COMP_CWORD]}
  //   local pre="${COMP_WORDS[COMP_CWORD - 1]}"
  //   # Pointer to current completion word.
  //   # By convention, it's named "cur" but this isn't strictly necessary.
  //
  //
  //   # echo-err "COMP_CWORD = ${COMP_CWORD}"
  //   # echo-err "cur= ${cur}"
  //
  //
  //   # TODO: Make a function that will check everything up to COMP_CWORD for the value that we're looking for
  //   # make it take an array of values or a space separated set of options to check for
  //   # then replace the mess below with a single function call
  //   #
  //   # if [ $(_previous_contains_all "${#COMP_WORDS[@]}" "-h --help help -V --version version") ]; then
  //   # if [ $(_previous_contains_some "${#COMP_WORDS[@]}" "-h --help help -V --version version") ]; then
  //   #
  //   # this will allow  random order of parameters
  //
  //   if [ "$pre" == "-h" -o "$pre" == "--help" -o "$pre" == "help" -o "$pre" == "-V" -o "$pre" == "--version" -o "$pre" == "version" ]; then
  //     return 0
  //   elif [ "$pre" == "clean" ]; then
  //     suggestions=($(compgen -W "--config --verbose"))
  //   elif [ "$pre" == "ls" -o "$pre" == "list" ]; then
  //     suggestions=($(compgen -W "--config --verbose -k --keywords-only"))
  //   else
  //     case "$cur" in
  //       -h | --help | help | -V | --version | version)
  //         return 0
  //         ;;
  //       --config)
  //         # suggestions should be local filenames
  //         ;;
  //       --*)
  //         suggestions=($(compgen -W "--help --config --verbose --version" -- "$cur"))
  //         ;;
  //       -*)
  //         suggestions=($(compgen -W "-h --help -V --version -v --verbose --config" -- "$cur"))
  //         ;;
  //       *)
  //         suggestions=($(compgen -W "-h --help -V --version -v --verbose --config clean ls list rm remove save help version $(curd ls -k)" -- "$cur"))
  //         ;;
  //     esac
  //   fi
  //
  //   COMPREPLY=("${suggestions[@]}")
  //
  //   return 0
}
