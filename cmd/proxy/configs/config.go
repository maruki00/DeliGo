package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Debug   string `yaml:"debug"`
	} `yaml:"app"`

	HTTPServer struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"http_server"`

	UserGRPC struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"user_grpc"`
}

var cfg *Config

func GetConfig() (*Config, error) {

	if cfg != nil {
		return cfg, nil
	}

	cfg = new(Config)
	conf, err := os.Open("./config.yaml")
	if err != nil {
		return nil, err
	}
	defer conf.Close()

	decoder := yaml.NewDecoder(conf)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
