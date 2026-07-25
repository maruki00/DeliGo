package pkgConfigs

import (
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

var (
	ErrOpenFile = errors.New("could not open config file")
	ErrDecode   = errors.New("could not decode config")
)

type App struct {
	Name     string `yaml:"name" env:"NAME"`
	Version  string `yaml:"version" env:"VERSION"`
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"`
}

func LoadCfg[T any](path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOpenFile, err)
	}
	defer file.Close()

	var cfg T
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDecode, err)
	}

	return &cfg, nil
}
