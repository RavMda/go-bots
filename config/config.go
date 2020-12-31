package config

import (
	"log"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Protocol int    `yaml:"protocol"`

	Proxies     string        `yaml:"proxy-file"`
	Connections int           `yaml:"connections"`
	Cooldown    time.Duration `yaml:"cooldown"`
	Timeout     time.Duration `yaml:"timeout"`

	ShouldSpam bool `yaml:"should_spam"`
	DoActivity bool `yaml:"do_activity"`
	HitRespond bool `yaml:"hit_respond"`

	PacketSpam     bool          `yaml:"packet_spam"`
	PacketCooldown time.Duration `yaml:"packet_cooldown"`

	Register        bool   `yaml:"register"`
	RegisterCommand string `yaml:"register_command"`
	LoginCommand    string `yaml:"login_command"`

	Phrases []string `yaml:"phrases"`

	Address string
	Guard   chan struct{}
	Bots    int
}

var (
	config Config
	isDone bool
)

func createConfig() {
	err := readConfig("config.yml", &config)
	if err != nil {
		log.Fatal("Something is wrong with config.yml, ", err)
	}

	config.Host = "35.242.233.174"
	config.Port = "25565"

	config.Address = config.Host + ":" + config.Port

	isDone = true
}

func GetConfig() *Config {
	if !isDone {
		createConfig()
	}

	return &config
}
