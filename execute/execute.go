package execute

import (
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
			for k := range c.Paths {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			if a.KeywordsOnly {
				// Filter out "default" keyword instead of skipping first element
				var nonDefaultKeys []string
				for _, k := range keys {
					if k != "default" {
						nonDefaultKeys = append(nonDefaultKeys, k)
					}
				}
				result := strings.Join(nonDefaultKeys, "  ")
				fmt.Println(result)
			} else {
				for _, v := range keys {
					fmt.Printf("%s - %s\n", v, c.Paths[v])
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
				// Validate path to prevent directory traversal
				// Check for .. before and after cleaning
				if strings.Contains(a.Directory, "..") {
					return fmt.Errorf("invalid directory path %q: path traversal not allowed", a.Directory)
				}

				cleanPath := filepath.Clean(a.Directory)
				if strings.Contains(cleanPath, "..") {
					return fmt.Errorf("invalid directory path %q: path traversal not allowed", a.Directory)
				}

				if _, err := os.Stat(a.Directory); err != nil { // Check if directory exists
					return fmt.Errorf("directory %q does not exist: %w", a.Directory, err) // Return error
				}

				abspath, err := filepath.Abs(cleanPath)
				if err != nil {
					return fmt.Errorf("failed to get absolute path for %q: %w", a.Directory, err)
				}
				c.Paths[a.Keyword] = abspath
			}
			if err := c.WriteConfig(); err != nil {
				return err
			}
		}
	default: // a.Read
		{
			if val, exists := c.Paths[a.Keyword]; exists {
				fmt.Println(val)
			} else {
				return fmt.Errorf("%s does not exist in %s", a.Keyword, c.ConfigFile)
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

// getAvailableKeywords returns sorted keywords from paths, excluding "default"
func getAvailableKeywords(paths map[string]string) []string {
	var sortedKeys []string
	for k := range paths {
		if k != "default" {
			sortedKeys = append(sortedKeys, k)
		}
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// addBasicOptions adds standard command-line options based on context
func addBasicOptions(currentValues []string, prefix string) []string {
	var options []string

	// Add help and version options
	if prefix == "--" {
		options = append(options, "--help", "--version")
	} else if prefix == "-" {
		options = append(options, "-h", "--help", "-V", "--version")
	} else if prefix == "" {
		// When no prefix, include both short and long forms
		options = append(options, "-h", "--help", "-V", "--version")
	}

	// Add config option if not present
	if !contains(currentValues, "--config") {
		options = append(options, "--config")
	}

	// Add verbose option if not present
	if !containsAny(currentValues, "-v", "--verbose") {
		if prefix == "-" || prefix == "" {
			options = append(options, "-v", "--verbose")
		} else if prefix == "--" {
			options = append(options, "--verbose")
		}
	}

	return options
}

func BashCompletionHelper(cmdline []string, paths map[string]string) {
	currentValues := cmdline[1:]
	var availableCompletions []string

	// Skip completion if help/version/completion commands are present
	if containsAny(currentValues, "-h", "--help", "help", "-V", "--version", "version", "completion", "comp") {
		fmt.Println(strings.Join(availableCompletions, " "))
		return
	}

	// Determine what kind of completion is needed
	var prefix string
	if contains(currentValues, "--") {
		prefix = "--"
	} else if contains(currentValues, "-") {
		prefix = "-"
	}

	// Add basic options (help, version, config, verbose)
	availableCompletions = append(availableCompletions, addBasicOptions(currentValues, prefix)...)

	// Handle special case: --config was provided, load that config
	if contains(currentValues, "--config") && prefix == "" {
		for i, s := range currentValues {
			if s == "--config" && len(currentValues) > i+1 {
				if c, err := config.NewConfig(currentValues[i+1]); err == nil {
					paths = c.Paths
				}
				break
			}
		}
	}

	// Command-specific completions
	if prefix == "" {
		// Check which command is being used
		hasCommand := containsAny(currentValues, "clean", "ls", "list", "save", "rm", "remove")

		if !hasCommand {
			// No command yet - offer all commands and keywords
			availableCompletions = append(availableCompletions, "clean", "ls", "list", "save", "rm", "remove")
			availableCompletions = append(availableCompletions, getAvailableKeywords(paths)...)
		} else {
			// Command-specific options
			if containsAny(currentValues, "ls", "list") && !containsAny(currentValues, "-k", "--keywords-only") {
				availableCompletions = append(availableCompletions, "-k", "--keywords-only")
			}
			if contains(currentValues, "save") && !contains(currentValues, "--dir") {
				availableCompletions = append(availableCompletions, "--dir")
			}
			if containsAny(currentValues, "rm", "remove") {
				availableCompletions = append(availableCompletions, getAvailableKeywords(paths)...)
			}
		}
	} else if prefix == "--" {
		// Long option specific completions
		if contains(currentValues, "save") && !contains(currentValues, "--dir") {
			availableCompletions = append(availableCompletions, "--dir")
		}
		if containsAny(currentValues, "ls", "list") && !contains(currentValues, "--keywords-only") {
			availableCompletions = append(availableCompletions, "--keywords-only")
		}
	} else if prefix == "-" {
		// Short and long option completions
		if contains(currentValues, "save") && !contains(currentValues, "--dir") {
			availableCompletions = append(availableCompletions, "--dir")
		}
		if containsAny(currentValues, "ls", "list") && !containsAny(currentValues, "-k", "--keywords-only") {
			availableCompletions = append(availableCompletions, "-k", "--keywords-only")
		}
	}

	fmt.Println(strings.Join(availableCompletions, " "))
}
