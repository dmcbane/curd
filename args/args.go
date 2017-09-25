package args

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"path"
	"sort"
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
    curd [KEYWORD] [-c FILE | --config=FILE] [--verbose]
    curd -n | --clean [-c FILE | --config=FILE] [--verbose]
    curd -l | --list [-c FILE | --config=FILE] [--verbose]
    curd [KEYWORD] -r | --remove [-c FILE | --config=FILE] [--verbose]
    curd [KEYWORD] -s | --save [-c FILE | --config=FILE] [--verbose]
    curd -h | --help
    curd -V | --version

Options:
    -c FILE, --config=FILE  Specify configuration filename [default: <<replaceme>>].
    -h, --help     Show this screen.
    -V, --version  Show version.
    -v, --verbose  Display extra information.

Examples:
    List all of the paths defined in the default configuration file.
        curd -l

    List all of the paths in a specified configuration file.
        curd --list --config some_configuration_file

    Clean paths from the default configuration that do not exist in the
    filesystem.
        curd --clean

    Clean paths that do not exist in the filesystem from the specified
    configuration.
        curd -n -c 'configuration file'

    Read the default path from the default configuration file.
        curd

    Save the current directory as the default path in the default configuration
    file.
        curd -s

    Remove the path associated with keyword 'essay' from the default
    configuration file.
        curd -r essay`

	usage = strings.Replace(usage, "<<replaceme>>", defaultConfig, 1)

	// parse the usage string
	// nil means use os.Args
	// true means display usage as the help message
	// the string to display for version
	// don't require options to be provided before positional arguments
	// have Parse call os.Exit() if help or version are requested by the user
	arguments, err := docopt.Parse(usage, nil, true, "Curd 1.0.0", false, false)
	if err != nil {
		fmt.Print(err)
	}
	var keys []string
	for k, _ := range arguments {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%9v: %v\n", k, arguments[k])
	}
	os.Exit(0)

	var cleanBool, listBool, readBool, removeBool, saveBool, verboseBool bool
	var keyword string

	configFile, _ = arguments["--config"].(string)
	directory, _ = arguments["--directory"].(string)
	keyword, _ = arguments["keyword"].(string)
	removeString, _ := arguments["--remove"]
	rString, _ := arguments["-r"]
	saveString, _ := arguments["--save"]
	sString, _ := arguments["-s"]
	cleanBool = arguments["--clean"].(bool) || arguments["-n"].(bool)
	listBool = arguments["--list"].(bool) || arguments["-l"].(bool)
	verboseBool = arguments["--verbose"].(bool)
	switch {
	case removeString != nil:
		removeBool = true
		keyword = removeString.(string)
	case rString != nil:
		removeBool = true
	case saveString != nil:
		saveBool = true
		keyword = saveString.(string)
	case sString != nil:
		saveBool = true
	}

	readBool = !cleanBool && !listBool && !removeBool && !saveBool

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
