package bot

import (
	"encoding/hex"
	"fmt"
	"go-pen/config"
	. "go-pen/config"
	. "go-pen/guard"
	"math/rand"
	"time"

	"net"

	"github.com/RavMda/go-mc/bot"
)

const (
	min = 4
	max = 15
)

func prepareBot(client *bot.Client, conn net.Conn, conf *config.Config) error {
	rand.Seed(time.Now().UnixNano())
	client.Auth.Name = GetName(rand.Intn(max-min+1) + min)

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client.JoinRaw(conn, conf.Address, conf.Protocol)
}

func destroyBot(data Data, reason string) {
	config := GetConfig()
	guard := GetGuard()

	fmt.Println("Bot left: ", reason)

	config.Bots--
	guard.Decrement()
}
