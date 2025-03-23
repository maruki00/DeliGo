package pkgConfigs

type AppCfg struct {
	Name     string `yaml:"name" env:"NAME"`
	Version  string `yaml:"version" env:"VERSION"`
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"`
}
