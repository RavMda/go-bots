package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host            string `yaml:"host" env:"HOST"`
	Port            string `yaml:"port" env:"PORT" env-default:"25565"`
	Proxies         string `yaml:"proxy-file" env:"PROXY_FILE" env-default:"proxies.txt"`
	RegisterCommand string `yaml:"register-command" env:"REGISTER_COMMAND" env-default:"register qweqwe123"`

	Connections int `yaml:"connections" env:"CONNECTIONS" env-default:"10"`
	Protocol    int `yaml:"protocol" env:"PROTOCOL" env-default:"754"`

	Register bool `yaml:"register" env:"REGISTER" env-default:"false"`

	Phrases    []string `yaml:"phrases"`
	ShouldSpam bool     `yaml:"should_spam" env:"SHOULD_SPAM" env-default:"false"`
	HitRespond bool     `yaml:"hit_respond" env:"HIT_RESPOND" env-default:"true"`

	Address string
	Guard   chan struct{}
}

var (
	config Config
	isDone bool
)

func createConfig() {
	err := cleanenv.ReadConfig("config.yml", &config)
	if err != nil {
		log.Fatal("Something is wrong with config.yml, ", err)
	}

	config.Host = "95.27.236.126"
	config.Port = "44444"

	isDone = true
}

func GetConfig() *Config {
	if !isDone {
		createConfig()
	}

	config.Host = "95.27.236.126"
	config.Port = "44444"

	return &config
}
