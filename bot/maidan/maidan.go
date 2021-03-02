package maidan

import (
	"fmt"
	"go-bots/bot"
	. "go-bots/config"
	"go-bots/methods"
	"math/rand"
	"net"
	"time"

	mcbot "github.com/RavMda/go-mc/bot"
	"github.com/RavMda/go-mc/bot/basic"
	"github.com/RavMda/go-mc/chat"
	"github.com/RavMda/go-mc/data/packetid"
	"github.com/RavMda/go-mc/net/packet"
	"github.com/google/uuid"
)

var config = GetConfig()

func CreateBot(conn net.Conn) {
	var client = mcbot.NewClient()

	config.Bots++

	if err := bot.PrepareBot(client, conn); err != nil {
		bot.DestroyBot(err.Error())
		return
	}

	basic.EventsListener{
		GameStart:  onGameStart,
		ChatMsg:    onChatMsg,
		Disconnect: onDisconnect,
		Death:        onDeath,
	}.Attach(client)

	err := client.HandleGame()
	if err != nil {
		fmt.Println("Handle Error: ", err)
	}
}

func sendMessage(client *mcbot.Client, message string) {
	client.Conn.WritePacket(
		packet.Marshal(
			packetid.ChatServerbound,
			packet.String(message),
		),
	)
}

func onGameStart(client *mcbot.Client) error {
	fmt.Println(config.Bots, "Bots connected")

	if config.Register {
		time.Sleep(10 * time.Second)

		sendMessage(client, config.RegisterCommand)
		time.Sleep(2 * time.Second)
		sendMessage(client, config.LoginCommand)
	}

	if config.ChatSpam {
		go doSpam(client)
	}

	if config.DoActivity {
		//go doActivity(client)
	}

	if config.PacketSpam {
		go sendPackets(client)
	}

	return nil
}

func onChatMsg(client *mcbot.Client, msg chat.Message, pos byte, uuid uuid.UUID) error {
	return nil
}

func onDisconnect(client *mcbot.Client, reason chat.Message) error {
	bot.DestroyBot(reason.String())
	return nil
}

func onDeath(client *mcbot.Client) error {
	return client.Conn.WritePacket(packet.Marshal(
		packetid.ClientCommand,
		packet.VarInt(0),
	))
}

func randomMessage(client *mcbot.Client) {
	phrases := config.Phrases
	phrase := phrases[rand.Intn(len(phrases))]

	sendMessage(client, phrase)
}

func doSpam(client *mcbot.Client) {
	for range time.Tick(config.ChatCooldown * time.Millisecond) {
		randomMessage(client)
	}
}

func sendPackets(client *mcbot.Client) {
	conn := client.Conn.Socket

	for range time.Tick(config.PacketCooldown * time.Millisecond) {
		if methods.Extreme1(conn) != nil {
			return
		}
	}
}

/*
func doActivity(client *mcbot.Client) {
	var num = float64(0)

	// arm swinging
	go func(client *mcbot.Client) {
		for range time.Tick(250 * time.Millisecond) {
			client.Conn.WritePacket(packet.Marshal(
				packetid.Animation,
				packet.VarInt(0),
			))
		}
	}(client)

	// client "movement"
	for range time.Tick(5 * time.Millisecond) {
		num = num + 0.1

		sin := math.Sin(num)
		cos := math.Cos(num)

		client.Conn.WritePacket(packetid.PositionAndLookServerbound{
			X:        packet.Double(client.Pos.X),
			Y:        packet.Double(client.Pos.Y + sin),
			Z:        packet.Double(client.Pos.Z),
			Yaw:      packet.Float(sin * 50),
			Pitch:    packet.Float(cos * 50),
			OnGround: packet.Boolean(sin < 0),
		}.Encode())

	}
}
*/
