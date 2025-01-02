package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host       string `toml:"host"`
	AuthCookie string `toml:"auth-cookie"`
	fileName   string `toml:"-"`
}

func (c Config) FileName() string {
	return c.fileName
}

func ReadConfig(filename string) (*Config, error) {
	var config Config
	md, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return nil, err
	}
	if len(md.Undecoded()) > 0 {
		return nil, fmt.Errorf("extra config keys: %v", md.Undecoded())
	}
	config.fileName = filename
	return &config, nil
}

func (c Config) Save() error {
	f, err := os.Create(c.FileName())
	if err != nil {
		return err
	}
	return toml.NewEncoder(f).Encode(c)
}
