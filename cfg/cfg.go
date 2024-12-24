package cfg

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host  string `yaml:"host"`
	Query string `yaml:"query"`
}

func MustLoad() *Config {
	configPath := fetchConfig()
	if configPath == "" {
		panic("empty config")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file doesn't exists " + configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Config is empty " + err.Error())
	}
	return &cfg
}
func fetchConfig() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res

}
