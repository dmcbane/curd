package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const cleanentrieskey = "clean"
const configfilekey = "configfile"
const entrykey = "entry"
const listentrykey = "list"
const removeentrykey = "remove"
const setentrykey = "set"

type byKey []string

func (k byKey) Len() int           { return len(k) }
func (k byKey) Less(i, j int) bool { return k[i] < k[j] }
func (k byKey) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

type Args struct {
	configFile string
	keyword    string
	clean      bool
	list       bool
	read       bool
	remove     bool
	save       bool
}

func GetCommandlineArguments() *Args {
	var configFile, defaultConfig, usage string

	defaultConfig = getDefaultConfigurationFilename()

	usage = fmt.Sprintf("Select a configuration file to use instead of the default (~/%s).", defaultConfig)
	flag.StringVar(&configFile, "c", "", usage)
	flag.StringVar(&configFile, "-config", "", usage)

	var cleanBool bool
	usage = "Cleanup entries for paths that don't exist."
	flag.BoolVar(&cleanBool, "n", false, usage)
	flag.BoolVar(&cleanBool, "-clean", false, usage)

	var listBool bool
	usage = "List all of the paths saved in the configuration file."
	flag.BoolVar(&listBool, "l", false, usage)
	flag.BoolVar(&listBool, "-list", false, usage)

	var removeBool bool
	usage = "Remove the path specified by the keyword or the default path from the configuration file."
	flag.BoolVar(&removeBool, "r", false, usage)
	flag.BoolVar(&removeBool, "-remove", false, usage)

	var saveBool bool
	usage = "Save the current directory to the specified keyword or the default."
	flag.BoolVar(&saveBool, "s", false, usage)
	flag.BoolVar(&saveBool, "-save", false, usage)
	flag.Parse()

	if configFile == "" {
		configFile = defaultConfig
	}

	readBool := !cleanBool && !listBool && !removeBool && !saveBool

	var keyword string
	values := flag.Args()
	if len(values) > 0 {
		keyword = flag.Args()[0]
	} else {
		keyword = "<default>"
	}
	return &Args{configFile: configFile, keyword: keyword, clean: cleanBool, list: listBool, read: readBool, remove: removeBool, save: saveBool}
}

func getDefaultConfigurationFilename() string {
	const default_config = ".curdrc"
	var home string
	if home = os.Getenv("HOME"); home == "" {
		home = os.Getenv("USERPROFILE")
	}
	return path.Join(home, default_config)
}
