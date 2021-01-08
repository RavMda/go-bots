package proxies

import (
	"go-bots/bot"
	"go-bots/config"
	"net"
	"time"

	"github.com/SteffenLoges/socks4"
)

func Dial(proxy string, address string) (net.Conn, bot.Dialer, error) {
	config := config.GetConfig()

	dialer := socks4.Dialer(socks4.SOCKS4, proxy, "", config.Timeout*time.Second)
	conn, err := dialer("tcp", address)

	return conn, dialer, err
}
