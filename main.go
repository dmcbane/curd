package main

import (
	"log"
	"os"

	"github.com/dmcbane/curd/args"
	"github.com/dmcbane/curd/config"
	"github.com/dmcbane/curd/execute"
)

var curdlog *log.Logger

func init() {
	// setup logger without date and time to be used
	// for reporting errors to stderr more than logging
	curdlog = log.New(os.Stderr, "curd: ", 0)
}

func main() {
	a := args.NewArgs()
	c, err := config.NewConfig(a.ConfigFile)
	if err != nil {
		curdlog.Fatalln(err)
	}
	if err = execute.ExecuteCommand(*a, *c); err != nil {
		curdlog.Fatalln(err)
	}
}
