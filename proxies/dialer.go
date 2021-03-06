package proxies

import (
	. "go-bots/config"
	"net"
	"time"

	"github.com/SteffenLoges/socks4"
)

func Dial(proxy string, address string) (net.Conn, error) {
	config := GetConfig()

	dialer := socks4.Dialer(socks4.SOCKS4, proxy, "", config.Timeout*time.Second)
	conn, err := dialer("tcp", address)

	return conn, err
}
