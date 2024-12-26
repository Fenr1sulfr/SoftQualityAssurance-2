package cfg

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host  string `yaml:"host"`
	Query string `yaml:"query"`
}

type LoginElements struct {
	UsernameElement string `yaml:"userElement"`
	PasswordElement string `yaml:"passwordElement"`
	ConfirmElement  string `yaml:"confirmElement"`
}

type TestLoginAndOut struct {
	URL           string `yaml:"url"`
	Login         string `yaml:"login"`
	Password      string `yaml:"password"`
	LoginElements `yaml:"elements"`
}

var (
	configLogin string
	configCase  string
)

func init() {
	// Define flags once in the init() function
	flag.StringVar(&configLogin, "config-login", "", "path to login config")
	flag.StringVar(&configCase, "config-case", "", "path to case config")
}

func MustLoadElements() *TestLoginAndOut {
	if configLogin == "" {
		configLogin = os.Getenv("CONFIG_PATH")
	}
	if configLogin == "" {
		panic("empty config")
	}
	if _, err := os.Stat(configLogin); os.IsNotExist(err) {
		panic("config file doesn't exist " + configLogin)
	}
	var tla TestLoginAndOut
	if err := cleanenv.ReadConfig(configLogin, &tla); err != nil {
		panic("Config is empty " + err.Error())
	}
	fmt.Println(tla)
	return &tla
}

func MustLoad() *Config {
	if configCase == "" {
		configCase = os.Getenv("CONFIG_PATH")
	}
	if configCase == "" {
		panic("empty config")
	}
	if _, err := os.Stat(configCase); os.IsNotExist(err) {
		panic("config file doesn't exist " + configCase)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configCase, &cfg); err != nil {
		panic("Config is empty " + err.Error())
	}
	return &cfg
}
