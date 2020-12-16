package bot

import (
	"fmt"
	"go-pen/config"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/chat"
	"github.com/RavMda/go-mc/net/packet"
	"github.com/RavMda/go-mc/net/ptypes"
)

func Maidan(conf *config.Config, conn net.Conn) {
	var client = bot.NewClient()

	if err := prepareBot(client, conn, conf); err != nil {
		destroyBot(client, err.Error())
	}

	client.Events.GameStart = onGameStart
	client.Events.Disconnect = onDisconnect
	client.Events.HealthChange = onHealthChange

	fmt.Println("Login success.")

	if err := client.HandleGame(); err != nil {
		log.Fatal(err)
	}
}

func onGameStart(client *bot.Client) error {
	log.Println("Game start", client.ID)

	if config.GetConfig().ShouldSpam {
		go doSpam(client)
	}

	go doActivity(client)

	return nil
}

func onDisconnect(client *bot.Client, reason chat.Message) error {
	destroyBot(client, "Bot left.")
	return nil
}

func onDeath(client *bot.Client) error {
	return client.Respawn()
}

func sendMessage(client *bot.Client) error {
	phrases := config.GetConfig().Phrases
	phrase := phrases[rand.Intn(len(phrases))]

	return client.Chat(phrase)
}

func onHealthChange(client *bot.Client, oldHealth float32, newHealth float32) error {
	if config.GetConfig().HitRespond && newHealth < oldHealth {
		sendMessage(client)
	}

	return nil
}

func doSpam(client *bot.Client) {
	config := config.GetConfig()

	if config.Register {
		client.Chat("/register qweqwe123")
		time.Sleep(2 * time.Second)
	}

	for {
		sendMessage(client)
		time.Sleep(1500 * time.Millisecond)
	}
}

func doActivity(client *bot.Client) {
	var num = new(float64)

	// arm swinging
	go func(client *bot.Client) {
		for {
			client.SwingArm(0)
			time.Sleep(250 * time.Millisecond)
		}
	}(client)

	// increment number for further usage
	go func(num *float64) {
		for {
			*num = *num + 0.1
			time.Sleep(10 * time.Millisecond)
		}
	}(num)

	// client "movement"
	for {
		sin := math.Sin(*num)
		cos := math.Cos(*num)

		client.Conn().WritePacket(ptypes.PositionAndLookServerbound{
			X:        packet.Double(client.Pos.X),
			Y:        packet.Double(client.Pos.Y + sin),
			Z:        packet.Double(client.Pos.Z),
			Yaw:      packet.Float(sin * 50),
			Pitch:    packet.Float(cos * 50),
			OnGround: packet.Boolean(sin < 0),
		}.Encode())

		time.Sleep(10 * time.Millisecond)
	}
}
