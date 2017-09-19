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

	"cmd/curd/internal/args"
	"cmd/curd/internal/config"
	"cmd/curd/internal/execute"
)

var curdlog *log.Logger

func init() {
	// setup logger without date and time to be used
	// for reporting errors to stderr more than logging
	curdlog = log.New(os.Stderr, "curd: ", 0)
}

func main() {
	args := args.GetCommandlineArguments()
	entries := config.ReadConfig((*args)[configfilekey])
	execute.ExecuteCommand(*args, *entries)
}
