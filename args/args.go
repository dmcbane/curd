package args

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"path"
	"strings"
)

type Args struct {
	ConfigFile string
	Keyword    string
	Clean      bool
	List       bool
	Read       bool
	Remove     bool
	Save       bool
	Verbose    bool
}

func NewArgs() *Args {
	var configFile, defaultConfig, directory, usage string

	defaultConfig = getDefaultConfigurationFilename()

	usage = `CURD - Change to a User's Recurring Directory 1.0.0
H. Dale McBane<h.dale.mcbane@gmail.com>
Save and return to paths you visit often.

Usage:
    curd clean [--config <file>] [--verbose]
    curd (ls | list) [--config <file>] [--verbose]
    curd remove [KEYWORD] [--config <file>] [--verbose]
    curd save [KEYWORD] [--dir <directory>] [--config <file>] [--verbose]
    curd [KEYWORD] [--config <file>] [--verbose]
    curd (-h | --help)
    curd (-V | --version)

Options:
    --config=<file>  Specify configuration filename [default: <<replaceme>>].
    --dir=<directory>  Specify configuration filename [default: <current directory>].
    -h, --help     Show this screen.
    -V, --version  Show version.
    -v, --verbose  Display extra information.

Examples:
    List all of the paths defined in the default configuration file.
        curd ls

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

`

	usage = strings.Replace(usage, "<<replaceme>>", defaultConfig, 1)

	// parse the usage string
	// nil means use os.Args
	// true means display usage as the help message
	// the string to display for version
	// don't require options to be provided before positional arguments
	// have Parse call os.Exit() if help or version are requested by the user
	arguments, _ := docopt.Parse(usage, nil, true, "Curd 1.0.0", false, true)

	var cleanBool, listBool, readBool, removeBool, saveBool, verboseBool bool
	var keyword string

	configFile, _ = arguments["--config"].(string)
	directory, _ = arguments["--dir"].(string)
	keyword, _ = arguments["KEYWORD"].(string)
	cleanBool = arguments["clean"].(bool)
	listBool = arguments["list"].(bool) || arguments["ls"].(bool)
	removeBool = arguments["remove"].(bool)
	saveBool = arguments["save"].(bool)
	readBool = !cleanBool && !listBool && !removeBool && !saveBool
	verboseBool = arguments["--verbose"].(bool)

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
		fmt.Printf("ConfigFile: %v\n", configFile)
		fmt.Printf("Keyword:    %v\n", keyword)
		fmt.Printf("Clean:      %v\n", cleanBool)
		fmt.Printf("List:       %v\n", listBool)
		fmt.Printf("Read:       %v\n", readBool)
		fmt.Printf("Remove:     %v\n", removeBool)
		fmt.Printf("Save:       %v\n", saveBool)
		fmt.Printf("Directory:  %v\n", directory)
		fmt.Printf("Verbose:    %v\n", verboseBool)
	}

	return &Args{ConfigFile: configFile, Keyword: keyword, Clean: cleanBool, List: listBool, Read: readBool, Remove: removeBool, Save: saveBool, Verbose: verboseBool}
}

func getDefaultConfigurationFilename() string {
	const default_config = ".curdrc"
	var home string
	if home = os.Getenv("HOME"); home == "" {
		home = os.Getenv("USERPROFILE")
	}
	return path.Join(home, default_config)
}
