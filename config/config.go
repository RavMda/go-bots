package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host            string `yaml:"host" env:"HOST"`
	Port            string `yaml:"port" env:"PORT" env-default:"25565"`
	Proxies         string `yaml:"proxy-file" env:"PROXY_FILE" env-default:"proxies.txt"`
	RegisterCommand string `yaml:"register_command" env:"REGISTER_COMMAND" env-default:"register qweqwe123"`
	LoginCommand    string `yaml:"login_command" env:"LOGIN_COMMAND" env-default:"login qweqwe123"`

	Connections int           `yaml:"connections" env:"CONNECTIONS" env-default:"10"`
	Cooldown    time.Duration `yaml:"cooldown" env:"COOLDOWN" env-default:"10"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5"`
	Protocol    int           `yaml:"protocol" env:"PROTOCOL" env-default:"754"`

	Register bool `yaml:"register" env:"REGISTER" env-default:"false"`

	Phrases    []string `yaml:"phrases"`
	DoActivity bool     `yaml:"do_activity" env:"DO_ACTIVITY" env-default:"false"`
	ShouldSpam bool     `yaml:"should_spam" env:"SHOULD_SPAM" env-default:"false"`
	HitRespond bool     `yaml:"hit_respond" env:"HIT_RESPOND" env-default:"true"`

	Address string
	Guard   chan struct{}
	Bots    int
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

	config.Host = "34.121.29.132"
	config.Port = "25565"

	isDone = true
}

func GetConfig() *Config {
	if !isDone {
		createConfig()
	}

	config.Host = "34.121.29.132"
	config.Port = "25565"

	return &config
}
