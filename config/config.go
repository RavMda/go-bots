package config

import (
	"log"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Protocol int    `yaml:"protocol"`

	Proxies      string        `yaml:"proxy-file"`
	Connections  int           `yaml:"connections"`
	SlowCooldown time.Duration `yaml:"slow_cooldown"`
	FastCooldown time.Duration `yaml:"fast_cooldown"`
	Timeout      time.Duration `yaml:"timeout"`

	ChatSpam     bool          `yaml:"chat_spam"`
	ChatCooldown time.Duration `yaml:"chat_cooldown"`
	DoActivity   bool          `yaml:"do_activity"`
	HitRespond   bool          `yaml:"hit_respond"`

	PacketSpam     bool          `yaml:"packet_spam"`
	PacketCooldown time.Duration `yaml:"packet_cooldown"`

	Register        bool   `yaml:"register"`
	RegisterCommand string `yaml:"register_command"`
	LoginCommand    string `yaml:"login_command"`

	Phrases []string `yaml:"phrases"`

	Address  string
	Guard    chan struct{}
	Bots     int
	Cooldown time.Duration
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

	config.Cooldown = config.FastCooldown
	config.Address = config.Host + ":" + config.Port

	isDone = true
}

func GetConfig() *Config {
	if !isDone {
		createConfig()
	}

	return &config
}
