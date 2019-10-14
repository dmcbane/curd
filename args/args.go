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
}

func NewArgs() *Args {
	var configFile, defaultConfig, directory, usage string

	defaultConfig = getDefaultConfigurationFilename()

	VERSION := "1.2.2"
	VERSION_USER := fmt.Sprintf("Curd %v", VERSION)
	usage = `CURD - Change to a User's Recurring Directory <<version>>
H. Dale McBane<h.dale.mcbane@gmail.com>
Save and return to paths you visit often.

Usage:
    curd clean [--config <file>] [--verbose]
    curd (completion | comp) CMDLINE ...
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

`

	usage = strings.Replace(usage, "<<version>>", VERSION, 1)
	usage = strings.Replace(usage, "<<replaceme>>", defaultConfig, 1)

	parser := &docopt.Parser{HelpHandler: docopt.PrintHelpAndExit, OptionsFirst: false, SkipHelpFlags: false}
	arguments, _ := parser.ParseArgs(usage, nil, VERSION_USER)

	var cleanBool, listBool, readBool, removeBool, saveBool, verboseBool, keywordsOnlyBool, completionBool bool
	var keyword string
        var cmdline []string

	if arguments["help"].(bool) {
		fmt.Println(usage)
		os.Exit(0)
	}
	if arguments["version"].(bool) {
		fmt.Println(VERSION_USER)
		os.Exit(0)
	}

	configFile, _ = arguments["--config"].(string)
	directory, _ = arguments["--dir"].(string)
	keywordsOnlyBool = arguments["--keywords-only"].(bool)
	keyword, _ = arguments["KEYWORD"].(string)
	cleanBool = arguments["clean"].(bool)
	listBool = arguments["list"].(bool) || arguments["ls"].(bool)
	removeBool = arguments["remove"].(bool) || arguments["rm"].(bool)
	saveBool = arguments["save"].(bool)
        completionBool = arguments["completion"].(bool) || arguments["comp"].(bool)
	readBool = !cleanBool && !listBool && !removeBool && !saveBool && !completionBool
	verboseBool = arguments["--verbose"].(bool)
        cmdline = arguments["CMDLINE"].([]string)


	if configFile == "" {
		configFile = defaultConfig
	}

	if directory == "<current directory>" {
		directory = ""
	}

	if keyword == "" {
		keyword = "<default>"
	}

	if verboseBool {
		fmt.Printf("ConfigFile:   %v\n", configFile)
		fmt.Printf("Keyword:      %v\n", keyword)
		fmt.Printf("Clean:        %v\n", cleanBool)
		fmt.Printf("List:         %v\n", listBool)
		fmt.Printf("Read:         %v\n", readBool)
		fmt.Printf("Remove:       %v\n", removeBool)
		fmt.Printf("Save:         %v\n", saveBool)
		fmt.Printf("Directory:    %v\n", directory)
		fmt.Printf("Verbose:      %v\n", verboseBool)
		fmt.Printf("KeywordsOnly: %v\n", keywordsOnlyBool)
		fmt.Printf("Completion:   %v\n", completionBool)
                fmt.Printf("Cmdline:      %v\n", cmdline)
	}

	return &Args{ConfigFile: configFile,
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
                Cmdline:      cmdline}
}

func getDefaultConfigurationFilename() string {
	const default_config = ".curdrc"
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}
	return filepath.Join(home, default_config)
}
