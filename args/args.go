package args

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Args struct {
	ConfigFile   string
	Keyword      string
	Clean        bool
	List         bool
	Read         bool
	Remove       bool
	Save         bool
	Directory    string
	Verbose      bool
	KeywordsOnly bool
	Completion   bool
	Cmdline      []string

	GenerateCompletions bool
	Shell               string
}

func getDefaultConfigurationFilename() string {
	const defaultConfig = ".curdrc"
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}
	if home == "" {
		// Fall back to current directory if home is not defined
		home = "."
		fmt.Fprintf(os.Stderr, "Warning: HOME/USERPROFILE not set, using current directory for config\n")
	}
	return filepath.Join(home, defaultConfig)
}

func generateUsage(version, defaultConfig string) string {
	usage := `CURD - Change to one of a User's Recurrent Directories <<version>>
H. Dale McBane<h.dale.mcbane@gmail.com>
Save and return to paths you visit often.

Usage:
    curd clean [--config <file>] [--verbose]
    curd (completion | comp) CMDLINE ...
    curd completions [<shell>]
    curd (ls | list) [-k | --keywords-only] [--config <file>] [--verbose]
    curd (rm | remove) [KEYWORD] [--config <file>] [--verbose]
    curd save [KEYWORD] [--dir <directory>] [--config <file>] [--verbose]
    curd (help | -h | --help)
    curd (version | -V | --version)
    curd [KEYWORD] [--config <file>] [--verbose]

Options:
    --config=<file>  Specify configuration filename [default: <<replaceme>>].
    --dir=<directory>  Specify path name to associate with keyword [default: <current directory>].
    -k, --keywords-only  Don't include the path names in the list command.
    -h, --help     Show this screen.
    -V, --version  Show version.
    -v, --verbose  Display extra information.

Examples:
    List all of the keywords and paths defined in the default configuration file.
        curd ls

    List all of the keywords defined in the default configuration file.
        curd ls -k

    List all of the paths in a specified configuration file.
        curd list --config some_configuration_file

    Clean paths from the default configuration that do not exist in the
    filesystem.
        curd clean

    Read the default path from the default configuration file.
        curd

    Save the current directory as the default path in the default configuration
    file.
        curd save

    Save the specified directory as the specified path in the default
    configuration file.
        curd save curd --dir=~/go/src/github.com/dmcbane/curd

    Remove the specified path from the default configuration file.
        curd remove essay

    Used by shell completion scripts.
        curd comp curd ls -

    Generate a shell completion script. SHELL may be bash, fish, or zsh; if
    omitted, the shell is detected from the SHELL environment variable.
        curd completions bash > ~/.curd_completion.bash

`

	usage = strings.Replace(usage, "<<version>>", version, 1)
	usage = strings.Replace(usage, "<<replaceme>>", defaultConfig, 1)
	return usage
}

func parseDocoptArgs(usage, versionUser string) (map[string]interface{}, error) {
	parser := &docopt.Parser{HelpHandler: docopt.PrintHelpAndExit, OptionsFirst: false, SkipHelpFlags: false}
	arguments, err := parser.ParseArgs(usage, nil, versionUser)
	if err != nil {
		return nil, err
	}

	if arguments["help"].(bool) {
		fmt.Println(usage)
		return nil, fmt.Errorf("help requested")
	}
	if arguments["version"].(bool) {
		fmt.Println(versionUser)
		return nil, fmt.Errorf("version requested")
	}
	return arguments, nil
}

func mapArguments(arguments map[string]interface{}, defaultConfig string) *Args {
	configFile, _ := arguments["--config"].(string)
	directory, _ := arguments["--dir"].(string)
	keywordsOnlyBool := arguments["--keywords-only"].(bool)
	keyword, _ := arguments["KEYWORD"].(string)
	cleanBool := arguments["clean"].(bool)
	listBool := arguments["list"].(bool) || arguments["ls"].(bool)
	removeBool := arguments["remove"].(bool) || arguments["rm"].(bool)
	saveBool := arguments["save"].(bool)
	completionBool := arguments["completion"].(bool) || arguments["comp"].(bool)
	generateCompletionsBool := arguments["completions"].(bool)
	shell, _ := arguments["<shell>"].(string)
	readBool := !cleanBool && !listBool && !removeBool && !saveBool && !completionBool && !generateCompletionsBool
	verboseBool := arguments["--verbose"].(bool)
	cmdline := arguments["CMDLINE"].([]string)

	if configFile == "" {
		configFile = defaultConfig
	}

	if directory == "<current directory>" {
		directory = ""
	}

	if keyword == "" {
		keyword = "default"
	}

	return &Args{
		ConfigFile:   configFile,
		Keyword:      keyword,
		Clean:        cleanBool,
		List:         listBool,
		Read:         readBool,
		Remove:       removeBool,
		Save:         saveBool,
		Directory:    directory,
		Verbose:      verboseBool,
		KeywordsOnly: keywordsOnlyBool,
		Completion:   completionBool,
		Cmdline:      cmdline,

		GenerateCompletions: generateCompletionsBool,
		Shell:               shell,
	}
}

func logVerbose(a *Args) {
	if a.Verbose {
		fmt.Printf("ConfigFile:   %v\n", a.ConfigFile)
		fmt.Printf("Keyword:      %v\n", a.Keyword)
		fmt.Printf("Clean:        %v\n", a.Clean)
		fmt.Printf("List:         %v\n", a.List)
		fmt.Printf("Read:         %v\n", a.Read)
		fmt.Printf("Remove:       %v\n", a.Remove)
		fmt.Printf("Save:         %v\n", a.Save)
		fmt.Printf("Directory:    %v\n", a.Directory)
		fmt.Printf("Verbose:      %v\n", a.Verbose)
		fmt.Printf("KeywordsOnly: %v\n", a.KeywordsOnly)
		fmt.Printf("Completion:   %v\n", a.Completion)
		fmt.Printf("Cmdline:      %v\n", a.Cmdline)
		fmt.Printf("GenerateCompletions: %v\n", a.GenerateCompletions)
		fmt.Printf("Shell:               %v\n", a.Shell)
	}
}

func NewArgs() *Args {
	defaultConfig := getDefaultConfigurationFilename()

	const VERSION = "2.2.0"
	versionUser := fmt.Sprintf("Curd %v", VERSION)
	usage := generateUsage(VERSION, defaultConfig)

	arguments, err := parseDocoptArgs(usage, versionUser)
	if err != nil {
		// This error is used to signal main to exit gracefully after printing help/version
		return nil
	}

	args := mapArguments(arguments, defaultConfig)
	logVerbose(args)
	return args
}
