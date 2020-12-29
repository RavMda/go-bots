package bot

import (
	"encoding/hex"
	"fmt"
	"go-pen/config"

	"net"

	"github.com/RavMda/go-mc/bot"
)

func prepareBot(client *bot.Client, conn net.Conn, conf *config.Config) error {
	client.Auth.Name = config.GetName()

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client.JoinRaw(conn, conf.Address, conf.Protocol)
}

func destroyBot(data Data, reason string) {
	config := config.GetConfig()

	fmt.Println("Bot left: ", reason)

	config.Bots--
	<-config.Guard
}
