package methods

import (
	"encoding/hex"
	"log"
	"net"
	"time"

	"github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/chat"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

var (
	client *bot.Client
	guard  chan struct{}
)

func prepareBot() *bot.Client {
	client = bot.NewClient()
	client.Auth.Name = "gamno_" + randstr.Hex(3)

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client
}

func joinServer(conn net.Conn, address string) error {
	return client.JoinRaw(conn, address, 753)
}

func CreateBot(data *Data, conn net.Conn, guard_n chan struct{}) {
	prepareBot()

	guard = guard_n

	err := joinServer(conn, data.Host+":"+data.Port)
	if err != nil {
		<-guard
		log.Println(err)
		return
	}

	log.Println("Login success")

	client.Events.ChatMsg = onChatMsg
	client.Events.GameStart = onGameStart
	client.Events.Die = onDeath
	client.Events.Disconnect = onDisconnect

	err = client.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onGameStart() error {
	log.Println("Game start")

	go func() {
		client.Chat("/register 123123q")
		time.Sleep(2 * time.Second)

		for {
			client.Chat("Owned by NoLogic")
			time.Sleep(1500 * time.Millisecond)
		}
	}()

	return nil
}

func onDisconnect(reason chat.Message) error {
	<-guard
	log.Println("Left: ", reason)
	return nil
}

func onChatMsg(msg chat.Message, pos byte, uuid uuid.UUID) error {
	log.Println("Chat:", msg.ClearString())
	return nil
}

func onDeath() error {
	client.Respawn()
	return nil
}
