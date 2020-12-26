package bot

import (
	"net"

	"github.com/RavMda/go-mc/bot"
)

type Data struct {
	Dialer func(string, string) (net.Conn, error)
	Client *bot.Client
	guard  chan struct{}
}
