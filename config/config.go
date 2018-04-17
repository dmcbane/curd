package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ConfigFile string
	Paths      map[string]string
}

func (c *Config) readConfig() error {
	c.Paths = make(map[string]string)

	if _, err := os.Stat(c.ConfigFile); err == nil {
		content, err := ioutil.ReadFile(c.ConfigFile)
		if err != nil {
			return err
		}
		m := make(map[interface{}]interface{})
		err = yaml.Unmarshal(content, &m)
		if err == nil {
			for k, v := range m {
				c.Paths[k.(string)] = v.(string)
			}
		}
	}
	return nil
}

func (c *Config) WriteConfig() error {
	content, err := yaml.Marshal(&(c.Paths))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.ConfigFile, content, 0644)
	if err != nil {
		return err
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
