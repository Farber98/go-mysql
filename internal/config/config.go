package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Database struct {
		User        string
		Password    string
		Host        string
		Db          string
		MaxOpenConn int
		MaxIdleConn int
		MaxTtlConn  int
	}
}

func InitConfig(configPath string) (*config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist")
	} else if err != nil {
		return nil, err
	}
	cfg := &config{}
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
