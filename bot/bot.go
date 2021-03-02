package bot

import (
	. "go-bots/config"
	. "go-bots/guard"
	"hash/fnv"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/offline"
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
		log.Fatal("Hash error: ", err)
	}

	return h.Sum64()
}

func makeSeed(address string) int64 {
	if config.ReuseName {
		return int64(hash(address))
	}

	return time.Now().UnixNano()
}

func PrepareBot(client *bot.Client, conn net.Conn) error {
	seed := makeSeed(conn.RemoteAddr().String())

	rand.Seed(seed)

	client.Auth.Name = GetName(rand.Intn(max-min+1)+min, seed)
	client.Auth.UUID = offline.NameToUUID(client.Auth.Name).String()

	return client.JoinRaw(conn, config.Address, config.Protocol)
}

func DestroyBot(reason string) {
	guard := GetGuard()

	//fmt.Printf("Bot %s left: %s\n", data.Client.Name, reason)

	config.Bots--
	guard.Decrement()
}
