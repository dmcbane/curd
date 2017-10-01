package execute

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/dmcbane/curd/args"
	"github.com/dmcbane/curd/config"
)

func ExecuteCommand(a args.Args, c config.Config) error {
	switch {
	case a.Clean:
		{
			for k, v := range c.Paths {
				if _, err := os.Stat(v); err != nil && os.IsNotExist(err) {
					delete(c.Paths, k)
				}
			}
			if err := c.WriteConfig(); err != nil {
				return err
			}
		}
	case a.List:
		{
			// sort the keys of the arguments map
			var keys []string
			for k, _ := range c.Paths {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, v := range keys {
				fmt.Printf("%s - %s\n", v, (c.Paths)[v])
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
				(c.Paths)[a.Keyword] = pwd
			} else {
				(c.Paths)[a.Keyword] = a.Directory
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
