package bot

import (
	"net"

	"github.com/RavMda/go-mc/bot"
)

type Dialer func(string, string) (net.Conn, error)

type Data struct {
	Dialer Dialer
	Client *bot.Client
}
