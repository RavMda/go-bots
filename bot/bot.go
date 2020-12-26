package bot

import (
	"encoding/hex"
	"fmt"
	"go-pen/config"

	"net"

	"github.com/RavMda/go-mc/bot"
	"github.com/thanhpk/randstr"
)

func prepareBot(client *bot.Client, conn net.Conn, config *config.Config) error {
	client.Auth.Name = "nologic_" + randstr.Hex(2)

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client.JoinRaw(conn, config.Address, config.Protocol)
}

func destroyBot(data Data, reason string) {
	fmt.Println("Bot left: ", reason)
	<-data.guard
}
