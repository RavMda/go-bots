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

func Basic(conn net.Conn, data Data) {
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

	fmt.Println("Connect success.")

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

	for range time.Tick(config.PacketCooldown * time.Millisecond) {
		if methods.Extreme1(conn) != nil {
			return
		}
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
	for range time.Tick(config.ChatCooldown * time.Millisecond) {
		sendMessage(client)
	}
}

func doActivity(client *bot.Client) {
	var num = float64(0)

	// arm swinging
	go func(client *bot.Client) {
		for range time.Tick(250 * time.Millisecond) {
			client.SwingArm(0)
		}
	}(client)

	// client "movement"
	for range time.Tick(5 * time.Millisecond) {
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

	}
}
