package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

const default_config = ".curdrc"
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

var curdlog *log.Logger

func init() {
	// setup logger without date and time to be used
	// for reporting errors to stderr more than logging
	curdlog = log.New(os.Stderr, "curd: ", 0)
}

func main() {
	args := getCommandlineArguments()
	entries := readConfig((*args)[configfilekey])

	switch {
	case (*args)[cleanentrieskey] == "true":
		{
			for k, v := range *entries {
				if _, err := os.Stat(v); err != nil && os.IsNotExist(err) {
					delete(*entries, k)
					writeConfig((*args)[configfilekey], entries)
				}
			}
		}
	case (*args)[listentrykey] == "true":
		{
			keys := make([]string, len(*entries))
			i := 0
			for k, _ := range *entries {
				keys[i] = k
				i++
			}
			sort.Sort(byKey(keys))
			for _, v := range keys {
				fmt.Printf("%s - %s\n", v, (*entries)[v])
			}
		}
	case (*args)[removeentrykey] == "true":
		{
			delete(*entries, (*args)[entrykey])
			writeConfig((*args)[configfilekey], entries)
		}
	case (*args)[setentrykey] == "true":
		{
			pwd, err := os.Getwd()
			if err != nil {
				curdlog.Fatalln(err)
			}
			(*entries)[(*args)[entrykey]] = pwd
			writeConfig((*args)[configfilekey], entries)
		}
	default:
		{
			if val, exists := (*entries)[(*args)[entrykey]]; exists {
				fmt.Println(val)
			} else {
				curdlog.Fatalf("%s does not exist in %s", (*args)[entrykey], (*args)[configfilekey])
			}
		}
	}
}

func getCommandlineArguments() *map[string]string {
	args := make(map[string]string)
	cleanBool := flag.Bool("n", false, "Cleanup entries for paths that don't exist.")
	configFile := flag.String("c", "", fmt.Sprintf("Select a configuration file to use instead of the default (~/%s).", default_config))
	listBool := flag.Bool("l", false, "List all of the paths saved in the configuration file.")
	removeBool := flag.Bool("r", false, "Remove the path specified by the keyword or the default path from the configuration file.")
	setBool := flag.Bool("s", false, "Save the current directory to the specified keyword or the default.")
	flag.Parse()
	args[cleanentrieskey] = fmt.Sprint(*cleanBool)
	args[configfilekey] = getConfigurationFilePath(*configFile)
	args[listentrykey] = fmt.Sprint(*listBool)
	args[removeentrykey] = fmt.Sprint(*removeBool)
	args[setentrykey] = fmt.Sprint(*setBool)
	values := flag.Args()
	if len(values) > 0 {
		args[entrykey] = flag.Args()[0]
	} else {
		args[entrykey] = "<default>"
	}
	return &args
}

func getConfigurationFilePath(fromUser string) string {
	if fromUser == "" {
		var home string
		if home = os.Getenv("HOME"); home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return path.Join(home, default_config)
	} else {
		return fromUser
	}
}

func readConfig(config_filename string) *map[string]string {
	configuration := make(map[string]string)

	file, err := os.Open(config_filename)
	defer file.Close()
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			ss := strings.Split(s, "|")
			configuration[strings.Trim(ss[0], " \t")] = strings.Trim(ss[1], " \t")
		}

		if err := scanner.Err(); err != nil {
			curdlog.Fatalln(err)
		}
	}
	return &configuration
}

func writeConfig(config_filename string, configuration *map[string]string) error {
	file, err := os.Create(config_filename)
	defer file.Close()
	if err != nil {
		curdlog.Fatalln(err)
	}
	for k, v := range *configuration {
		_, err := file.WriteString(fmt.Sprintf("%s| %s\n", k, v))
		if err != nil {
			curdlog.Fatalln(err)
		}
	}
	return nil
}
