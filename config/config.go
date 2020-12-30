package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Proxies         string `yaml:"proxy-file"`
	RegisterCommand string `yaml:"register_command"`
	LoginCommand    string `yaml:"login_command"`

	Connections int           `yaml:"connections"`
	Cooldown    time.Duration `yaml:"cooldown"`
	Timeout     time.Duration `yaml:"timeout"`
	Protocol    int           `yaml:"protocol"`

	Register   bool `yaml:"register"`
	DoActivity bool `yaml:"do_activity"`
	ShouldSpam bool `yaml:"should_spam"`
	HitRespond bool `yaml:"hit_respond"`

	PacketSpam     bool          `yaml:"packet_spam"`
	PacketCooldown time.Duration `yaml:"packet_cooldown"`

	Address string
	Guard   chan struct{}
	Bots    int

	Phrases []string `yaml:"phrases"`
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
