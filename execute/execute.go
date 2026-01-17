package execute

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
				if k == "default" { // Explicitly skip cleaning the "default" entry
					continue
				}
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
				c.Paths[a.Keyword] = pwd
			} else {
				if _, err := os.Stat(a.Directory); err != nil { // Check if directory exists
					return fmt.Errorf("directory %q does not exist: %w", a.Directory, err) // Return error
				}
				abspath, _ := filepath.Abs(filepath.Clean(a.Directory))
				c.Paths[a.Keyword] = abspath
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

// Helper functions for slice operations
func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func containsAny(slice []string, checks ...string) bool {
	for _, value := range checks {
		if contains(slice, value) {
			return true
		}
	}
	return false
}

func BashCompletionHelper(cmdline []string, paths map[string]string) {
	currentValues := cmdline[1:]
	var availableCompletions []string
	// drop the first value since it is always curd

	// completions are only available if not one of these
	if !containsAny(currentValues, "-h", "--help", "help", "-V", "--version", "version", "completion", "comp") {
		if contains(currentValues, "--") { // starting a long option
			availableCompletions = append(availableCompletions, "--help")
			availableCompletions = append(availableCompletions, "--version")
			if !contains(currentValues, "--config") {
				availableCompletions = append(availableCompletions, "--config")
			}
			if !contains(currentValues, "--verbose") {
				availableCompletions = append(availableCompletions, "--verbose")
			}
			if contains(currentValues, "save") && !contains(currentValues, "--dir") {
				availableCompletions = append(availableCompletions, "--dir")
			}

			if containsAny(currentValues, "ls", "list") {
				if !containsAny(currentValues, "-k", "--keywords-only") {
					availableCompletions = append(availableCompletions, "--keywords-only")
				}
			}
		} else if contains(currentValues, "-") { // starting a long or short options
			availableCompletions = append(availableCompletions, "-h")
			availableCompletions = append(availableCompletions, "--help")
			availableCompletions = append(availableCompletions, "-V")
			availableCompletions = append(availableCompletions, "--version")
			if !contains(currentValues, "--config") {
				availableCompletions = append(availableCompletions, "--config")
			}
			if !containsAny(currentValues, "-v", "--verbose") {
				availableCompletions = append(availableCompletions, "-v")
				availableCompletions = append(availableCompletions, "--verbose")
			}
			if contains(currentValues, "save") && !contains(currentValues, "--dir") {
				availableCompletions = append(availableCompletions, "--dir")
			}

			if containsAny(currentValues, "ls", "list") {
				if !containsAny(currentValues, "-k", "--keywords-only") {
					availableCompletions = append(availableCompletions, "-k", "--keywords-only")
				}
			}

		} else {
			availableCompletions = append(availableCompletions, "-h")
			availableCompletions = append(availableCompletions, "--help")
			availableCompletions = append(availableCompletions, "-V")
			availableCompletions = append(availableCompletions, "--version")
			if !contains(currentValues, "--config") {
				availableCompletions = append(availableCompletions, "--config")
			} else {
				var configFile string
				for i, s := range currentValues {
					if s == "--config" {
						if len(currentValues) > i+1 {
							configFile = currentValues[i+1]
						}
						break
					}
				}
				c, err := config.NewConfig(configFile)
				if err == nil {
					paths = c.Paths
				}
			}

			if !containsAny(currentValues, "-v", "--verbose") {
				availableCompletions = append(availableCompletions, "-v")
				availableCompletions = append(availableCompletions, "--verbose")
			}
			// if none of the command group exists
			if !containsAny(currentValues, "clean", "ls", "list", "save", "rm", "remove") {
				// add all commands to the completions list
				availableCompletions = append(availableCompletions, "clean", "ls", "list", "save", "rm", "remove")
				// add all defined paths to the completions list
				var sortedKeys []string
				for k := range paths {
					if k != "default" {
						sortedKeys = append(sortedKeys, k)
					}
				}
				sort.Strings(sortedKeys)
				availableCompletions = append(availableCompletions, sortedKeys...)
			} else { // at least one command exists, so let's find out which
				// if contains(currentValues, "clean") {
				//   // nothing to do for clean
				// }

				if containsAny(currentValues, "ls", "list") {
					if !containsAny(currentValues, "-k", "--keywords-only") {
						availableCompletions = append(availableCompletions, "-k", "--keywords-only")
					}
				}

				if contains(currentValues, "save") {
					if !contains(currentValues, "--dir") {
						availableCompletions = append(availableCompletions, "--dir")
					}
				}

				if containsAny(currentValues, "rm", "remove") {
					// add all defined paths to the completions list
					var sortedKeys []string
					for k := range paths {
						if k != "default" {
							sortedKeys = append(sortedKeys, k)
						}
					}
					sort.Strings(sortedKeys)
					availableCompletions = append(availableCompletions, sortedKeys...)
				}
			}
		}
	}
	fmt.Println(strings.Join(availableCompletions, " "))
}
