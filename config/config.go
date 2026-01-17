package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ConfigFile string
	Paths      map[string]string
}

func (c *Config) readConfig() error {
	c.Paths = make(map[string]string)

	if _, err := os.Stat(c.ConfigFile); err == nil {
		content, err := os.ReadFile(c.ConfigFile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, &c.Paths)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) WriteConfig() error {
	content, err := yaml.Marshal(&(c.Paths))
	if err != nil {
		return err
	}
	err = os.WriteFile(c.ConfigFile, content, 0644)
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
