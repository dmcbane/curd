package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	ConfigFile string
	Paths      map[string]string
}

func (c *Config) readConfig() error {
	c.Paths = make(map[string]string)

	file, err := os.Open(c.ConfigFile)
	defer file.Close()
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			ss := strings.Split(s, "|")
			c.Paths[strings.Trim(ss[0], " \t")] = strings.Trim(ss[1], " \t")
		}
		return scanner.Err()
	} else {
		return err
	}
}

func (c *Config) WriteConfig() error {
	file, err := os.Create(c.ConfigFile)
	defer file.Close()
	if err != nil {
		return err
	}
	for k, v := range c.Paths {
		if _, err := file.WriteString(fmt.Sprintf("%s| %s\n", k, v)); err != nil {
			return err
		}
	}
	return nil
}

func NewConfig(filename string) (*Config, error) {
	c := Config{ConfigFile: filename}
	if err := c.readConfig(); err == nil {
		return &c, nil
	} else {
		return nil, err
	}
}
