package bot

import (
	"fmt"
	"go-bots/methods"

	"log"
	"math"
	"math/rand"
	"net"

	"time"

	"github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/chat"
	_ "github.com/RavMda/go-mc/data/lang/en-us"
	_ "github.com/RavMda/go-mc/data/lang/ru-ru"
	"github.com/RavMda/go-mc/net/packet"
	"github.com/RavMda/go-mc/net/ptypes"
)

func Maidan(conn net.Conn, data Data) {
	var client = bot.NewClient()

	data.Client = client

	config.Bots++

	if err := prepareBot(client, conn, config); err != nil {
		destroyBot(data, err.Error())
		return
	}

	client.Events.GameStart = onGameStart
	client.Events.Disconnect = onDisconnect
	client.Events.HealthChange = onHealthChange
	client.Events.Die = onDeath

	fmt.Println("Login success.")

	if err := client.HandleGame(); err != nil {
		log.Fatal(err)
	}
}

func onGameStart(client *bot.Client) error {
	fmt.Println(config.Bots, "Bots connected")

	if config.Register {
		time.Sleep(10 * time.Second)
		client.Chat(config.RegisterCommand)
		time.Sleep(2 * time.Second)
		client.Chat(config.LoginCommand)
		time.Sleep(2 * time.Second)
	}

	if config.ChatSpam {
		go doSpam(client)
	}

	if config.DoActivity {
		go doActivity(client)
	}

	if config.PacketSpam {
		go sendPackets(client)
	}

	return nil
}

func sendPackets(client *bot.Client) {
	conn := client.Conn().Socket
	cooldown := config.PacketCooldown * time.Millisecond

	for {
		if methods.Extreme1(conn) != nil {
			return
		}

		time.Sleep(cooldown)
	}
}

func onDisconnect(client *bot.Client, reason chat.Message) error {
	destroyBot(Data{Client: client}, reason.String())
	return nil
}

func onDeath(client *bot.Client) error {
	return client.Respawn()
}

func sendMessage(client *bot.Client) error {
	phrases := config.Phrases
	phrase := phrases[rand.Intn(len(phrases))]

	return client.Chat(phrase)
}

func onHealthChange(client *bot.Client, oldHealth float32, newHealth float32) error {
	if config.HitRespond && newHealth < oldHealth {
		sendMessage(client)
	}

	return nil
}

func doSpam(client *bot.Client) {
	for {
		sendMessage(client)
		time.Sleep(2000 * time.Millisecond)
	}
}

func doActivity(client *bot.Client) {
	var num = float64(0)

	// arm swinging
	go func(client *bot.Client) {
		for {
			client.SwingArm(0)
			time.Sleep(250 * time.Millisecond)
		}
	}(client)

	// client "movement"
	for {
		num = num + 0.1

		sin := math.Sin(num)
		cos := math.Cos(num)

		client.Conn().WritePacket(ptypes.PositionAndLookServerbound{
			X:        packet.Double(client.Pos.X),
			Y:        packet.Double(client.Pos.Y + sin),
			Z:        packet.Double(client.Pos.Z),
			Yaw:      packet.Float(sin * 50),
			Pitch:    packet.Float(cos * 50),
			OnGround: packet.Boolean(sin < 0),
		}.Encode())

		time.Sleep(5 * time.Millisecond)
	}
}
