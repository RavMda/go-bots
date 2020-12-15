package methods

import (
	"encoding/hex"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/chat"
	"github.com/RavMda/go-mc/net/packet"
	"github.com/RavMda/go-mc/net/ptypes"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

var (
	guard chan struct{}
)

func prepareBot() *bot.Client {
	client := bot.NewClient()
	client.Auth.Name = "nologic_" + randstr.Hex(2)

	id := bot.OfflineUUID(client.Auth.Name)
	client.Auth.UUID = hex.EncodeToString(id[:])

	return client
}

func joinServer(client *bot.Client, conn net.Conn, address string) error {
	return client.JoinRaw(conn, address, 754)
}

func CreateBot(data *Data, conn net.Conn, guard_n chan struct{}) {
	var client = prepareBot()

	guard = guard_n

	err := joinServer(client, conn, data.Host+":"+data.Port)
	if err != nil {
		<-guard
		log.Println(err)
		return
	}

	log.Println("Login success")

	client.Events.ChatMsg = onChatMsg
	client.Events.GameReady = onGameReady
	client.Events.Disconnect = onDisconnect

	err = client.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onGameReady(client *bot.Client) error {
	log.Println("Game start")

	go doSpam(client)
	go doActivity(client)
	go doJump(client)

	return nil
}

func doActivity(client *bot.Client) {
	for {
		client.SwingArm(0)
		time.Sleep(250 * time.Millisecond)
	}
}

func doJump(client *bot.Client) {
	var num = new(float64)

	go func(num *float64) {
		for {
			*num = *num + 0.1
			time.Sleep(10 * time.Millisecond)
		}
	}(num)

	for {
		add := math.Sin(*num)

		client.Conn().WritePacket(ptypes.Position{
			X:        packet.Double(client.Pos.X),
			Y:        packet.Double(client.Pos.Y + add),
			Z:        packet.Double(client.Pos.Z),
			OnGround: packet.Boolean(add < 0),
		}.Encode())

		time.Sleep(10 * time.Millisecond)
	}
}

var phrases = []string{
	"phrase",
}

func doSpam(client *bot.Client) {
	//client.Chat("/register qweqwe123")
	//time.Sleep(2 * time.Second)

	for {
		client.Chat(phrases[rand.Intn(len(phrases))])
		time.Sleep(1500 * time.Millisecond)
	}
}

func onDisconnect(reason chat.Message) error {
	<-guard
	log.Println("Left: ", reason.Text)
	return nil
}

func onChatMsg(msg chat.Message, pos byte, uuid uuid.UUID) error {
	//log.Println("Chat:", msg.ClearString())
	return nil
}
