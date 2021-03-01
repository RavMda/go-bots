package bot

import (
	"encoding/hex"
	"fmt"
	. "go-bots/config"
	. "go-bots/guard"
	"hash/fnv"
	"log"
	"math/rand"
	"time"

	"net"

	"github.com/RavMda/go-mc/bot"
)

const (
	min = 4
	max = 15
)

var config = GetConfig()

func hash(s string) uint64 {
	h := fnv.New64a()

	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	return h.Sum64()
}

func makeSeed(address string) int64 {
	if config.ReuseName {
		return int64(hash(address))
	}

	return time.Now().UnixNano()
}

func prepareBot(client *bot.Client, conn net.Conn, conf *Config) error {
	seed := makeSeed(conn.RemoteAddr().String())

	rand.Seed(seed)
	client.Auth.Name = GetName(rand.Intn(max-min+1)+min, seed)

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client.JoinRaw(conn, conf.Address, conf.Protocol)
}

func destroyBot(data Data, reason string) {
	guard := GetGuard()

	fmt.Printf("Bot %s left: %s\n", data.Client.Name, reason)

	config.Bots--
	guard.Decrement()
}
