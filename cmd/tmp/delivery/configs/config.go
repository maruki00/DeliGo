package shared_configs

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Debug   string `yaml:"debug"`
	} `yaml:"app"`

	RestServer struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"rest_server"`

	GRPCServer struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"grpc_server"`

	Postgres struct {
		Dsn string `yaml:"ddsn"`
	} `yaml:"postgres"`

	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	RabbitMq struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"rabbitmq"`

	Queues []string `yaml:"queues"`
}

var cfg *Config

func GetConfig(root string) (*Config, error) {

	if cfg != nil {
		return cfg, nil
	}

	cfg = new(Config)
	conf, err := os.Open(path.Join(root, "config.yaml"))
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
