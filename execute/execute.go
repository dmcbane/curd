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
			BashCompletionHelper(a.Cmdline, c.Paths)
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

type StringArray struct {
	value []string
}

func _extend(slice []string, item string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, n, 2*n+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = item
	return slice
}

func _append(slice []string, items ...string) []string {
	for _, item := range items {
		slice = _extend(slice, item)
	}
	return slice
}

func (a StringArray) contains(check string) bool {
	for _, value := range a.value {
		if value == check {
			return true
		}
	}
	return false
}

func (a StringArray) containsAll(checks ...string) bool {
	for _, value := range checks {
		if !a.contains(value) {
			return false
		}
	}
	return true
}

func (a StringArray) containsAny(checks ...string) bool {
	for _, value := range checks {
		if a.contains(value) {
			return true
		}
	}
	return false
}

func (a StringArray) toString() string {
	return strings.Join(a.value, " ")
}

func BashCompletionHelper(cmdline []string, paths map[string]string) {
	var currentValues = new(StringArray)
	var availableCompletions = new(StringArray)
	currentValues.toString()
	availableCompletions.toString()

	currentValues.value = cmdline[1:]
	// drop the first value since it is always curd
	currentValues.toString()

	// completions are only available if not one of these
	if !currentValues.containsAny("-h", "--help", "help", "-V", "--version", "version", "completion", "comp") {
		if !currentValues.contains("--config") {
			availableCompletions.value = _append(availableCompletions.value, "--config")
			availableCompletions.toString()
		} else {
			var configFile string
			for i, s := range currentValues.value {
				if s == "--config" {
					configFile = currentValues.value[i+1]
					break
				}
			}
			c, err := config.NewConfig(configFile)
			if err == nil {
				paths = c.Paths
			}
		}

		if !currentValues.contains("--verbose") {
			availableCompletions.value = _append(availableCompletions.value, "--verbose")
			availableCompletions.toString()
		}
		// if none of the command group exists
		if !currentValues.containsAny("clean", "ls", "list", "save", "rm", "remove") {
			// add all commands to the completions list
			availableCompletions.value = _append(availableCompletions.value, "clean", "ls", "list", "save", "rm", "remove")
			// add all defined paths to the completions list
			for k, _ := range paths {
				if k != "<default>" {
					availableCompletions.value = _append(availableCompletions.value, k)
				}
			}
		} else { // at least one command exists, so let's find out which
			// if currentValues.contains("clean") {
			//   // nothing to do for clean
			// }

			if currentValues.containsAny("ls", "list") {
				if !currentValues.containsAny("-k", "--keywords-only") {
					availableCompletions.value = _append(availableCompletions.value, "-k", "--keywords-only")
				}
			}

			if currentValues.containsAny("save") {
				if !currentValues.containsAny("--dir") {
					availableCompletions.value = _append(availableCompletions.value, "--dir")
				}
			}

			if currentValues.containsAny("rm", "remove") {
				// add all defined paths to the completions list
				for k, _ := range paths {
					if k != "<default>" {
						availableCompletions.value = _append(availableCompletions.value, k)
					}
				}
			}
		}
	}
	fmt.Println(availableCompletions.toString())
}
